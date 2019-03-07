package api_test

import (
	"errors"
	"fmt"
	"net/http"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/configuration/coreconfig"
	"code.cloudfoundry.org/cli/cf/net"
	"code.cloudfoundry.org/cli/cf/terminal/terminalfakes"
	"code.cloudfoundry.org/cli/cf/trace/tracefakes"
	testconfig "code.cloudfoundry.org/cli/cf/util/testhelpers/configuration"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/gomega/ghttp"
)

var _ = Describe("ClientRepository", func() {
	FDescribe("ClientExists", func() {
		var (
			client     api.ClientRepository
			uaaServer  *ghttp.Server
			uaaGateway net.Gateway
			config     coreconfig.ReadWriter
		)

		BeforeEach(func() {
			uaaServer = ghttp.NewServer()
			config = testconfig.NewRepositoryWithDefaults()

			config.SetUaaEndpoint(uaaServer.URL())
			uaaGateway = net.NewUAAGateway(config, new(terminalfakes.FakeUI), new(tracefakes.FakePrinter), "")
			client = api.NewCloudControllerClientRepository(config, uaaGateway)
		})

		Context("when a client does not exist", func() {
			var clientID string
			BeforeEach(func() {
				clientID = "some-client"

				requestPath := fmt.Sprintf("/oauth/clients/%s", clientID)

				uaaServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", requestPath),
						ghttp.RespondWith(http.StatusNotFound, ""),
					),
				)
			})

			It("returns an error", func() {

				b, err := client.ClientExists("some-client")
				Expect(err).To(Not(BeNil()))
				Expect(b).To(BeFalse())
			})
		})

		Context("when a client does exist", func() {
			var clientID string
			BeforeEach(func() {
				clientID = "some-client"

				requestPath := fmt.Sprintf("/oauth/clients/%s", clientID)

				uaaServer.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", requestPath),
						ghttp.RespondWith(http.StatusOK, ""),
					),
				)
			})
			It("returns true and no error", func() {
				b, err := client.ClientExists("some-client")

				Expect(b).To(BeTrue())
				Expect(err).To(BeNil())

			})
		})

		Context("when getAuthEndpoint fails", func() {
			var err error

			BeforeEach(func() {
				err = errors.New("auth-endpoint-error")
				client.getAuthEndpointReturns("bad-endpoint-path", err)
			})

			It("returns that error", func() {
				b, err := client.ClientExists("some-client")
			})
		})
	})

})
