// +build windows,!windows2012R2

package handlers_test

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"code.cloudfoundry.org/diego-release/diego-ssh/daemon"
	"code.cloudfoundry.org/diego-release/diego-ssh/handlers"
	"code.cloudfoundry.org/diego-release/diego-ssh/handlers/fakes"
	"code.cloudfoundry.org/diego-release/diego-ssh/test_helpers"
	"code.cloudfoundry.org/lager/lagertest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var _ = Describe("SessionChannelHandler", func() {
	type command struct {
		path string
		args []string
	}

	var (
		sshd   *daemon.Daemon
		client *ssh.Client

		commandsRan chan command

		logger          *lagertest.TestLogger
		serverSSHConfig *ssh.ServerConfig

		runner                *fakes.FakeRunner
		shellLocator          *fakes.FakeShellLocator
		sessionChannelHandler *handlers.SessionChannelHandler

		newChannelHandlers map[string]handlers.NewChannelHandler
		defaultEnv         map[string]string
		connectionFinished chan struct{}
	)

	BeforeEach(func() {
		logger = lagertest.NewTestLogger("test")
		serverSSHConfig = &ssh.ServerConfig{
			NoClientAuth: true,
		}
		serverSSHConfig.AddHostKey(TestHostKey)

		commandsRan = make(chan command, 10)

		runner = &fakes.FakeRunner{}
		realRunner := handlers.NewCommandRunner()
		runner.StartStub = func(cmd *exec.Cmd) error {
			commandsRan <- command{
				path: strings.ToLower(cmd.Path),
				args: cmd.Args,
			}
			return realRunner.Start(cmd)
		}
		runner.WaitStub = realRunner.Wait
		runner.SignalStub = realRunner.Signal

		shellLocator = &fakes.FakeShellLocator{}
		shellLocator.ShellPathReturns("cmd.exe")

		defaultEnv = map[string]string{}
		for _, env := range os.Environ() {
			k := strings.Split(env, "=")[0]
			v := strings.Split(env, "=")[1]
			defaultEnv[k] = v
		}
		defaultEnv["TEST"] = "FOO"

		delete(defaultEnv, "Path")
		delete(defaultEnv, "PATH")

		sessionChannelHandler = handlers.NewSessionChannelHandler(runner, shellLocator, defaultEnv, time.Second)

		newChannelHandlers = map[string]handlers.NewChannelHandler{
			"session": sessionChannelHandler,
		}

		serverNetConn, clientNetConn := test_helpers.Pipe()

		sshd = daemon.New(logger, serverSSHConfig, nil, newChannelHandlers)
		connectionFinished = make(chan struct{})
		go func() {
			sshd.HandleConnection(serverNetConn)
			close(connectionFinished)
		}()

		client = test_helpers.NewClient(clientNetConn, nil)
	})

	AfterEach(func() {
		if client != nil {
			err := client.Close()
			Expect(err).NotTo(HaveOccurred())
		}
		Eventually(connectionFinished).Should(BeClosed())
	})

	Context("when a session is opened", func() {
		var session *ssh.Session

		BeforeEach(func() {
			var sessionErr error
			session, sessionErr = client.NewSession()

			Expect(sessionErr).NotTo(HaveOccurred())
		})

		It("can use the session to execute a command with stdout and stderr", func() {
			stdout, err := session.StdoutPipe()
			Expect(err).NotTo(HaveOccurred())

			stderr, err := session.StderrPipe()
			Expect(err).NotTo(HaveOccurred())

			err = session.Run("echo Hello && echo Goodbye 1>&2")
			Expect(err).NotTo(HaveOccurred())

			stdoutBytes, err := ioutil.ReadAll(stdout)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(stdoutBytes)).To(ContainSubstring("Hello"))
			Expect(string(stdoutBytes)).NotTo(ContainSubstring("Goodbye"))

			stderrBytes, err := ioutil.ReadAll(stderr)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(stderrBytes)).To(ContainSubstring("Goodbye"))
			Expect(string(stderrBytes)).NotTo(ContainSubstring("Hello"))
		})

		It("returns when the process exits", func() {
			stdin, err := session.StdinPipe()
			Expect(err).NotTo(HaveOccurred())

			err = session.Run("dir")
			Expect(err).NotTo(HaveOccurred())

			stdin.Close()
		})

		Describe("scp", func() {
			var (
				sourceDir, generatedTextFile, targetDir string
				err                                     error
				stdin                                   io.WriteCloser
				stdout                                  io.Reader
				fileContents                            []byte
			)

			BeforeEach(func() {
				stdin, err = session.StdinPipe()
				Expect(err).NotTo(HaveOccurred())

				stdout, err = session.StdoutPipe()
				Expect(err).NotTo(HaveOccurred())

				sourceDir, err = ioutil.TempDir("", "scp-source")
				Expect(err).NotTo(HaveOccurred())

				fileContents = []byte("---\nthis is a simple file\n\n")
				generatedTextFile = filepath.Join(sourceDir, "textfile.txt")

				err = ioutil.WriteFile(generatedTextFile, fileContents, 0664)
				Expect(err).NotTo(HaveOccurred())

				targetDir, err = ioutil.TempDir("", "scp-target")
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				Expect(os.RemoveAll(sourceDir)).To(Succeed())
				Expect(os.RemoveAll(targetDir)).To(Succeed())
			})

			It("properly copies using the secure copier", func() {
				done := make(chan struct{})
				go func() {
					defer GinkgoRecover()
					err := session.Run(fmt.Sprintf("scp -v -t %s", strings.Replace(targetDir, `\`, `\\`, -1)))
					Expect(err).NotTo(HaveOccurred())
					close(done)
				}()

				confirmation := make([]byte, 1)
				_, err = stdout.Read(confirmation)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmation).To(Equal([]byte{0}))

				expectedFileInfo, err := os.Stat(generatedTextFile)
				Expect(err).NotTo(HaveOccurred())

				_, err = stdin.Write([]byte(fmt.Sprintf("C0664 %d textfile.txt\n", expectedFileInfo.Size())))
				Expect(err).NotTo(HaveOccurred())

				_, err = stdout.Read(confirmation)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmation).To(Equal([]byte{0}))

				_, err = stdin.Write(fileContents)
				Expect(err).NotTo(HaveOccurred())

				_, err = stdin.Write([]byte{0})
				Expect(err).NotTo(HaveOccurred())

				_, err = stdout.Read(confirmation)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmation).To(Equal([]byte{0}))

				err = stdin.Close()
				Expect(err).NotTo(HaveOccurred())

				actualFilePath := filepath.Join(targetDir, filepath.Base(generatedTextFile))
				actualFileInfo, err := os.Stat(actualFilePath)
				Expect(err).NotTo(HaveOccurred())

				Expect(actualFileInfo.Mode()).To(Equal(expectedFileInfo.Mode()))
				Expect(actualFileInfo.Size()).To(Equal(expectedFileInfo.Size()))

				actualContents, err := ioutil.ReadFile(actualFilePath)
				Expect(err).NotTo(HaveOccurred())

				expectedContents, err := ioutil.ReadFile(generatedTextFile)
				Expect(err).NotTo(HaveOccurred())

				Expect(actualContents).To(Equal(expectedContents))

				Eventually(done).Should(BeClosed())
			})

			It("properly fails when secure copying fails", func() {
				errCh := make(chan error)
				go func() {
					defer GinkgoRecover()
					errCh <- session.Run(fmt.Sprintf("scp -v -t %s", strings.Replace(targetDir, `\`, `\\`, -1)))
				}()

				confirmation := make([]byte, 1)
				_, err = stdout.Read(confirmation)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmation).To(Equal([]byte{0}))

				_, err = stdin.Write([]byte("BOGUS PROTOCOL MESSAGE\n"))
				Expect(err).NotTo(HaveOccurred())

				_, err = stdout.Read(confirmation)
				Expect(err).NotTo(HaveOccurred())
				Expect(confirmation).To(Equal([]byte{1}))

				err = <-errCh
				exitErr, ok := err.(*ssh.ExitError)
				Expect(ok).To(BeTrue())
				Expect(exitErr.ExitStatus()).To(Equal(1))
			})

			It("properly fails when incorrect arguments are supplied", func() {
				err := session.Run(fmt.Sprintf("scp -v -t /tmp/foo /tmp/bar"))
				Expect(err).To(HaveOccurred())

				exitErr, ok := err.(*ssh.ExitError)
				Expect(ok).To(BeTrue())
				Expect(exitErr.ExitStatus()).To(Equal(1))
			})
		})

		Describe("the shell locator", func() {
			BeforeEach(func() {
				err := session.Run("exit 0")
				Expect(err).NotTo(HaveOccurred())
			})

			It("uses the shell locator to find the default shell path", func() {
				Expect(shellLocator.ShellPathCallCount()).To(Equal(1))

				Eventually(commandsRan).Should(Receive(Equal(command{
					path: "c:\\windows\\system32\\cmd.exe",
					args: []string{"cmd.exe", "/c", "exit 0"},
				})))
			})
		})

		Context("when stdin is provided by the client", func() {
			BeforeEach(func() {
				session.Stdin = strings.NewReader("Hello")
			})

			It("can use the session to execute a command that reads it", func() {
				result, err := session.Output(`findstr x*`)

				Expect(err).NotTo(HaveOccurred())
				Expect(strings.TrimSpace(string(result))).To(Equal("Hello"))
			})
		})

		Context("when the command exits with a non-zero value", func() {
			It("it preserve the exit code", func() {
				err := session.Run("exit 3")
				Expect(err).To(HaveOccurred())

				exitErr, ok := err.(*ssh.ExitError)
				Expect(ok).To(BeTrue())
				Expect(exitErr.ExitStatus()).To(Equal(3))
			})
		})

		Context("when SIGKILL is sent across the session", func() {
			Context("before a command has been run", func() {
				BeforeEach(func() {
					err := session.Signal(ssh.SIGKILL)
					Expect(err).NotTo(HaveOccurred())
				})

				It("does not prevent the command from running", func() {
					result, err := session.Output("echo still kicking")
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(result))).To(Equal(strings.TrimSpace("still kicking")))
				})
			})

			Context("while a command is running", func() {
				var stdin io.WriteCloser
				var stdout io.Reader

				BeforeEach(func() {
					var err error
					stdin, err = session.StdinPipe()
					Expect(err).NotTo(HaveOccurred())

					stdout, err = session.StdoutPipe()
					Expect(err).NotTo(HaveOccurred())

					err = session.Shell()
					Expect(err).NotTo(HaveOccurred())

					reader := bufio.NewReader(stdout)
					Eventually(reader.ReadLine).Should(ContainSubstring("Microsoft Windows"))

					Eventually(runner.StartCallCount).Should(Equal(1))
				})

				It("is sent to the process", func() {
					err := session.Signal(ssh.SIGKILL)
					Expect(err).NotTo(HaveOccurred())

					Eventually(runner.SignalCallCount).Should(Equal(1))

					err = stdin.Close()
					if err != nil {
						Expect(err).To(Equal(io.EOF), "expected no error or ignorable EOF error")
					}

					err = session.Wait()
					Expect(err).To(HaveOccurred())
					Expect(err.(*ssh.ExitError).ExitStatus()).To(Equal(1))
				})
			})
		})

		Context("when running a command without an explicit environemnt", func() {
			It("does not inherit daemon's environment", func() {
				os.Setenv("DAEMON_ENV", "daemon_env_value")

				result, err := session.Output("set")
				Expect(err).NotTo(HaveOccurred())

				Expect(result).NotTo(ContainSubstring("DAEMON_ENV=daemon_env_value"))
				os.Unsetenv("DAEMON_ENV")
			})

			It("includes a default environment excluding PATH", func() {
				result, err := session.Output("set")
				Expect(err).NotTo(HaveOccurred())

				Expect(string(result)).To(ContainSubstring(fmt.Sprintf(`PATH=C:\Windows\system32;C:\Windows;C:\Windows\System32\Wbem;C:\Windows\System32\WindowsPowerShell\v1.0`)))
				Expect(result).To(ContainSubstring(fmt.Sprintf("LANG=en_US.UTF8")))
				Expect(result).To(ContainSubstring(fmt.Sprintf("TEST=FOO")))
				Expect(result).To(ContainSubstring(fmt.Sprintf("HOME=%s", os.Getenv("HOME"))))
				Expect(result).To(ContainSubstring(fmt.Sprintf("USER=%s", os.Getenv("USER"))))
			})
		})

		Context("when environment variables are requested", func() {
			Context("before starting the command", func() {
				It("runs the command with the specified environment", func() {
					err := session.Setenv("ENV1", "value1")
					Expect(err).NotTo(HaveOccurred())

					err = session.Setenv("ENV2", "value2")
					Expect(err).NotTo(HaveOccurred())

					result, err := session.Output("set")
					Expect(err).NotTo(HaveOccurred())

					Expect(result).To(ContainSubstring("ENV1=value1"))
					Expect(result).To(ContainSubstring("ENV2=value2"))
				})

				It("uses the value last specified", func() {
					err := session.Setenv("ENV1", "original")
					Expect(err).NotTo(HaveOccurred())

					err = session.Setenv("ENV1", "updated")
					Expect(err).NotTo(HaveOccurred())

					result, err := session.Output("set")
					Expect(err).NotTo(HaveOccurred())

					Expect(result).To(ContainSubstring("ENV1=updated"))
				})

				It("can override PATH and LANG", func() {
					err := session.Setenv("PATH", "/bin:/usr/local/bin:/sbin")
					Expect(err).NotTo(HaveOccurred())

					err = session.Setenv("LANG", "en_UK.UTF8")
					Expect(err).NotTo(HaveOccurred())

					result, err := session.Output("set")
					Expect(err).NotTo(HaveOccurred())

					Expect(result).To(ContainSubstring("PATH=/bin:/usr/local/bin:/sbin"))
					Expect(result).To(ContainSubstring("LANG=en_UK.UTF8"))
				})

				It("cannot override HOME and USER", func() {
					err := session.Setenv("HOME", "/some/other/home")
					Expect(err).NotTo(HaveOccurred())

					err = session.Setenv("USER", "not-a-user")
					Expect(err).NotTo(HaveOccurred())

					result, err := session.Output("set")
					Expect(err).NotTo(HaveOccurred())

					Expect(result).To(ContainSubstring(fmt.Sprintf("HOME=%s", os.Getenv("HOME"))))
					Expect(result).To(ContainSubstring(fmt.Sprintf("USER=%s", os.Getenv("USER"))))
				})

				It("can override default env variables", func() {
					err := session.Setenv("TEST", "BAR")
					Expect(err).NotTo(HaveOccurred())

					result, err := session.Output("set")
					Expect(err).NotTo(HaveOccurred())

					Expect(result).To(ContainSubstring("TEST=BAR"))
				})
			})

			Context("after starting the command", func() {
				var stdin io.WriteCloser
				var stdout io.Reader

				BeforeEach(func() {
					var err error
					stdin, err = session.StdinPipe()
					Expect(err).NotTo(HaveOccurred())

					stdout, err = session.StdoutPipe()
					Expect(err).NotTo(HaveOccurred())

					err = session.Start(`findstr x* & set`)
					Expect(err).NotTo(HaveOccurred())
				})

				It("ignores the request", func() {
					err := session.Setenv("ENV3", "value3")
					Expect(err).NotTo(HaveOccurred())

					stdin.Close()

					err = session.Wait()
					Expect(err).NotTo(HaveOccurred())

					stdoutBytes, err := ioutil.ReadAll(stdout)
					Expect(err).NotTo(HaveOccurred())

					Expect(string(stdoutBytes)).NotTo(ContainSubstring("ENV3"))
				})
			})
		})

		Context("when a pty request is received", func() {
			var terminalModes ssh.TerminalModes

			BeforeEach(func() {
				terminalModes = ssh.TerminalModes{}
			})

			JustBeforeEach(func() {
				err := session.RequestPty("vt100", 43, 80, terminalModes)
				Expect(err).NotTo(HaveOccurred())
			})

			It("should allocate a console for the session", func() {
				result, err := session.Output("timeout 1 2>nul >nul & if errorlevel 1 (echo redirect) else (echo console)")
				Expect(err).NotTo(HaveOccurred())

				Expect(string(result)).To(ContainSubstring("console"))
			})

			It("returns when the process exits", func() {
				stdin, err := session.StdinPipe()
				Expect(err).NotTo(HaveOccurred())

				err = session.Run("dir")
				Expect(err).NotTo(HaveOccurred())

				stdin.Close()
			})

			It("terminates the shell when the stdin closes", func() {
				err := session.Shell()
				Expect(err).NotTo(HaveOccurred())
				time.Sleep(1 * time.Second)

				err = client.Conn.Close()
				client = nil
				Expect(err).NotTo(HaveOccurred())
				err = session.Wait()
				Expect(err.Error()).To(Equal("wait: remote command exited without exit status or exit signal"))
			})

			It("should set the terminal type", func() {
				result, err := session.Output(`echo %TERM%`)
				Expect(err).NotTo(HaveOccurred())

				Expect(string(result)).To(ContainSubstring("vt100"))
			})

			It("sets the correct window size for the terminal", func() {
				result, err := session.Output("powershell.exe -command $w = $host.ui.rawui.WindowSize.Width; $h = $host.ui.rawui.WindowSize.Height; echo \"$h $w\"")
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(ContainSubstring("43 80"))
			})

			Context("when an interactive command is executed", func() {
				var stdin io.WriteCloser

				JustBeforeEach(func() {
					var err error
					stdin, err = session.StdinPipe()
					Expect(err).NotTo(HaveOccurred())
				})

				It("terminates the session when the shell exits", func() {
					err := session.Start("cmd.exe")
					Expect(err).NotTo(HaveOccurred())

					_, err = stdin.Write([]byte("exit\r\n"))
					Expect(err).NotTo(HaveOccurred())

					err = stdin.Close()
					Expect(err).NotTo(HaveOccurred())

					err = session.Wait()
					Expect(err).NotTo(HaveOccurred())
				})
			})

			Context("when a signal is sent across the session", func() {
				Context("before a command has been run", func() {
					BeforeEach(func() {
						err := session.Signal(ssh.SIGKILL)
						Expect(err).NotTo(HaveOccurred())
					})

					It("does not prevent the command from running", func() {
						result, err := session.Output("echo still kicking")
						Expect(err).NotTo(HaveOccurred())
						Expect(string(result)).To(ContainSubstring("still kicking"))
					})
				})

				Context("SIGKILL is sent while a command is running", func() {
					var stdin io.WriteCloser
					var stdout io.Reader

					JustBeforeEach(func() {
						var err error
						stdin, err = session.StdinPipe()
						Expect(err).NotTo(HaveOccurred())

						stdout, err = session.StdoutPipe()
						Expect(err).NotTo(HaveOccurred())

						err = session.Shell()
						Expect(err).NotTo(HaveOccurred())

						reader := bufio.NewReader(stdout)
						Eventually(reader.ReadLine).Should(ContainSubstring("Microsoft Windows"))
					})

					It("kills the process", func() {
						err := session.Signal(ssh.SIGKILL)
						Expect(err).NotTo(HaveOccurred())

						err = stdin.Close()
						if err != nil {
							Expect(err).To(Equal(io.EOF), "expected no error or ignorable EOF error")
						}

						err = session.Wait()
						Expect(err).To(HaveOccurred())
						Expect(err.(*ssh.ExitError).ExitStatus()).To(Equal(1))
					})
				})

				Context("SIGINT is sent while a command is running", func() {
					var stdin io.WriteCloser
					var stdout io.Reader

					JustBeforeEach(func() {
						var err error
						stdin, err = session.StdinPipe()
						Expect(err).NotTo(HaveOccurred())

						stdout, err = session.StdoutPipe()
						Expect(err).NotTo(HaveOccurred())

						err = session.Start("ping -t 127.0.0.1 & echo goodbye")
						Expect(err).NotTo(HaveOccurred())

						reader := bufio.NewReader(stdout)
						Eventually(reader.ReadLine).Should(ContainSubstring("127.0.0.1"))
					})

					It("the process is interrupted", func() {
						err := session.Signal(ssh.SIGINT)
						Expect(err).NotTo(HaveOccurred())

						err = stdin.Close()
						if err != nil {
							Expect(err).To(Equal(io.EOF), "expected no error or ignorable EOF error")
						}

						resultCh := make(chan error)
						go func() {
							resultCh <- session.Wait()
						}()
						Eventually(resultCh).Should(Receive(BeNil()))
						reader := bufio.NewReader(stdout)
						Eventually(reader.ReadLine).Should(ContainSubstring("goodbye"))
					})
				})
			})
		})

		Context("when a window change request is received", func() {
			type winChangeMsg struct {
				Columns  uint32
				Rows     uint32
				WidthPx  uint32
				HeightPx uint32
			}

			var result []byte

			Context("before a pty is allocated", func() {
				BeforeEach(func() {
					_, err := session.SendRequest("window-change", false, ssh.Marshal(winChangeMsg{
						Rows:    50,
						Columns: 132,
					}))
					Expect(err).NotTo(HaveOccurred())

					err = session.RequestPty("vt100", 43, 80, ssh.TerminalModes{})
					Expect(err).NotTo(HaveOccurred())

					result, err = session.Output("powershell.exe -command $w = $host.ui.rawui.WindowSize.Width; $h = $host.ui.rawui.WindowSize.Height; echo \"$h $w\"")
					Expect(err).NotTo(HaveOccurred())
				})

				It("ignores the request", func() {
					Expect(result).To(ContainSubstring("43 80"))
				})
			})

			Context("after a pty is allocated", func() {
				BeforeEach(func() {
					err := session.RequestPty("vt100", 43, 80, ssh.TerminalModes{})
					Expect(err).NotTo(HaveOccurred())

					_, err = session.SendRequest("window-change", false, ssh.Marshal(winChangeMsg{
						Rows:    50,
						Columns: 132,
					}))
					Expect(err).NotTo(HaveOccurred())

					result, err = session.Output("powershell.exe -command $w = $host.ui.rawui.WindowSize.Width; $h = $host.ui.rawui.WindowSize.Height; echo \"$h $w\"")
					Expect(err).NotTo(HaveOccurred())
				})

				It("changes the the size of the terminal", func() {
					Expect(result).To(ContainSubstring("50 132"))
				})
			})
		})

		Context("after executing a command", func() {
			BeforeEach(func() {
				err := session.Run("exit")
				Expect(err).NotTo(HaveOccurred())
			})

			It("the session is no longer usable", func() {
				_, err := session.SendRequest("exec", true, ssh.Marshal(struct{ Command string }{Command: "exit"}))
				Expect(err).To(HaveOccurred())

				_, err = session.SendRequest("bogus", true, nil)
				Expect(err).To(HaveOccurred())

				err = session.Setenv("foo", "bar")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when an interactive shell is requested", func() {
			var stdin io.WriteCloser

			BeforeEach(func() {
				var err error
				stdin, err = session.StdinPipe()
				Expect(err).NotTo(HaveOccurred())

				err = session.Shell()
				Expect(err).NotTo(HaveOccurred())
			})

			It("starts the shell with the runner", func() {
				Eventually(runner.StartCallCount).Should(Equal(1))

				Eventually(commandsRan).Should(Receive(Equal(command{
					path: "c:\\windows\\system32\\cmd.exe",
					args: []string{"cmd.exe"},
				})))
			})

			It("terminates the session when the shell exits", func() {
				_, err := stdin.Write([]byte("exit\n"))
				Expect(err).NotTo(HaveOccurred())

				err = stdin.Close()
				if err != nil {
					Expect(err).To(Equal(io.EOF), "expected no error or ignorable EOF error")
				}

				err = session.Wait()
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("and a command is provided", func() {
			BeforeEach(func() {
				err := session.Run("exit")
				Expect(err).NotTo(HaveOccurred())
			})

			It("uses the provided runner to start the command", func() {
				Expect(runner.StartCallCount()).To(Equal(1))
				Expect(runner.WaitCallCount()).To(Equal(1))
			})

			It("passes the correct command to the runner", func() {
				Eventually(commandsRan).Should(Receive(Equal(command{
					path: "c:\\windows\\system32\\cmd.exe",
					args: []string{"cmd.exe", "/c", "exit"},
				})))
			})

			It("passes the same command to Start and Wait", func() {
				command := runner.StartArgsForCall(0)
				Expect(runner.WaitArgsForCall(0)).To(Equal(command))
			})
		})

		Context("when executing an invalid command", func() {
			It("returns an exit error with a non-zero exit status", func() {
				err := session.Run("not-a-command")
				Expect(err).To(HaveOccurred())

				exitErr, ok := err.(*ssh.ExitError)
				Expect(ok).To(BeTrue())
				Expect(exitErr.ExitStatus()).NotTo(Equal(0))
			})

			Context("when starting the command fails", func() {
				BeforeEach(func() {
					runner.StartReturns(errors.New("oops"))
				})

				It("returns an exit status message with a non-zero status", func() {
					err := session.Run("true")
					Expect(err).To(HaveOccurred())

					exitErr, ok := err.(*ssh.ExitError)
					Expect(ok).To(BeTrue())
					Expect(exitErr.ExitStatus()).NotTo(Equal(0))
				})
			})

			Context("when waiting on the command fails", func() {
				BeforeEach(func() {
					runner.WaitReturns(errors.New("oops"))
				})

				It("returns an exit status message with a non-zero status", func() {
					err := session.Run("true")
					Expect(err).To(HaveOccurred())

					exitErr, ok := err.(*ssh.ExitError)
					Expect(ok).To(BeTrue())
					Expect(exitErr.ExitStatus()).NotTo(Equal(0))
				})
			})
		})

		Context("when an unknown request type is sent", func() {
			var accepted bool

			BeforeEach(func() {
				var err error
				accepted, err = session.SendRequest("unknown-request-type", true, []byte("payload"))
				Expect(err).NotTo(HaveOccurred())
			})

			It("rejects the request", func() {
				Expect(accepted).To(BeFalse())
			})

			It("does not terminate the session", func() {
				response, err := session.Output("echo Hello")
				Expect(err).NotTo(HaveOccurred())
				Expect(strings.TrimSpace(string(response))).To(Equal("Hello"))
			})
		})

		Context("when an unknown subsystem is requested", func() {
			var accepted bool

			BeforeEach(func() {
				type subsysMsg struct{ Subsystem string }

				var err error
				accepted, err = session.SendRequest("subsystem", true, ssh.Marshal(subsysMsg{Subsystem: "unknown"}))
				Expect(err).NotTo(HaveOccurred())
			})

			It("rejects the request", func() {
				Expect(accepted).To(BeFalse())
			})

			It("does not terminate the session", func() {
				response, err := session.Output("echo Hello")
				Expect(err).NotTo(HaveOccurred())
				Expect(strings.TrimSpace(string(response))).To(Equal("Hello"))
			})
		})
	})

	Context("when the sftp subystem is requested", func() {
		It("accepts the request", func() {
			type subsysMsg struct{ Subsystem string }
			session, err := client.NewSession()
			Expect(err).NotTo(HaveOccurred())
			defer session.Close()

			accepted, err := session.SendRequest("subsystem", true, ssh.Marshal(subsysMsg{Subsystem: "sftp"}))
			Expect(err).NotTo(HaveOccurred())
			Expect(accepted).To(BeTrue())
		})

		It("starts an sftp server in write mode", func() {
			tempDir, err := ioutil.TempDir("", "sftp")
			Expect(err).NotTo(HaveOccurred())
			defer os.RemoveAll(tempDir)

			sftp, err := sftp.NewClient(client)
			Expect(err).NotTo(HaveOccurred())
			defer sftp.Close()

			By("creating the file")
			target := filepath.Join(tempDir, "textfile.txt")
			file, err := sftp.Create(target)
			Expect(err).NotTo(HaveOccurred())

			fileContents := []byte("---\nthis is a simple file\n\n")
			_, err = file.Write(fileContents)
			Expect(err).NotTo(HaveOccurred())

			err = file.Close()
			Expect(err).NotTo(HaveOccurred())

			Expect(ioutil.ReadFile(target)).To(Equal(fileContents))

			By("reading the file")
			file, err = sftp.Open(target)
			Expect(err).NotTo(HaveOccurred())

			buffer := &bytes.Buffer{}
			_, err = buffer.ReadFrom(file)
			Expect(err).NotTo(HaveOccurred())

			err = file.Close()
			Expect(err).NotTo(HaveOccurred())

			Expect(buffer.Bytes()).To(Equal(fileContents))

			By("removing the file")
			err = sftp.Remove(target)
			Expect(err).NotTo(HaveOccurred())

			_, err = os.Stat(target)
			Expect(err).To(HaveOccurred())
			Expect(os.IsNotExist(err)).To(BeTrue())
		})
	})

	Describe("invalid session channel requests", func() {
		var channel ssh.Channel
		var requests <-chan *ssh.Request

		BeforeEach(func() {
			var err error
			channel, requests, err = client.OpenChannel("session", nil)
			Expect(err).NotTo(HaveOccurred())

			go ssh.DiscardRequests(requests)
		})

		AfterEach(func() {
			if channel != nil {
				channel.Close()
			}
		})

		Context("when an exec request fails to unmarshal", func() {
			It("rejects the request", func() {
				accepted, err := channel.SendRequest("exec", true, ssh.Marshal(struct{ Bogus uint32 }{Bogus: 1138}))
				Expect(err).NotTo(HaveOccurred())
				Expect(accepted).To(BeFalse())
			})
		})

		Context("when an env request fails to unmarshal", func() {
			It("rejects the request", func() {
				accepted, err := channel.SendRequest("env", true, ssh.Marshal(struct{ Bogus int }{Bogus: 1234}))
				Expect(err).NotTo(HaveOccurred())
				Expect(accepted).To(BeFalse())
			})
		})

		Context("when a signal request fails to unmarshal", func() {
			It("rejects the request", func() {
				accepted, err := channel.SendRequest("signal", true, ssh.Marshal(struct{ Bogus int }{Bogus: 1234}))
				Expect(err).NotTo(HaveOccurred())
				Expect(accepted).To(BeFalse())
			})
		})

		Context("when a pty request fails to unmarshal", func() {
			It("rejects the request", func() {
				accepted, err := channel.SendRequest("pty-req", true, ssh.Marshal(struct{ Bogus int }{Bogus: 1234}))
				Expect(err).NotTo(HaveOccurred())
				Expect(accepted).To(BeFalse())
			})
		})

		Context("when a window change request fails to unmarshal", func() {
			It("rejects the request", func() {
				accepted, err := channel.SendRequest("window-change", true, ssh.Marshal(struct{ Bogus int }{Bogus: 1234}))
				Expect(err).NotTo(HaveOccurred())
				Expect(accepted).To(BeFalse())
			})
		})

		Context("when a subsystem request fails to unmarshal", func() {
			It("rejects the request", func() {
				accepted, err := channel.SendRequest("subsystem", true, ssh.Marshal(struct{ Bogus int }{Bogus: 1234}))
				Expect(err).NotTo(HaveOccurred())
				Expect(accepted).To(BeFalse())
			})
		})
	})
})
