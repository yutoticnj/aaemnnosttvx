package noop

import (
	"fmt"

	sdkModels "github.com/edgexfoundry/device-sdk-go/v2/pkg/models"
	"github.com/edgexfoundry/go-mod-core-contracts/v2/models"

	"github.com/edgexfoundry/device-camera-go/internal/pkg/client"
)

// Client is a camera client for cameras which don't have any other camera or manufacturer
// specific clients to leverage.  All of this client's methods return an error if information is
// requested and otherwise comply silently with direction to Initialize or Release.
type Client struct{}

// HandleReadCommand triggers a protocol Read operation for the specified device, resulting in
// an error for an unrecognized read command.
func (n Client) HandleReadCommand(req sdkModels.CommandRequest) (*sdkModels.CommandValue, error) {
	return &sdkModels.CommandValue{}, fmt.Errorf("device-camera-go: unrecognized read command")
}

// HandleWriteCommand triggers a protocol Write operation; resulting in an error for an unrecognized write command
func (n Client) HandleWriteCommand(req sdkModels.CommandRequest, param *sdkModels.CommandValue) error {
	return fmt.Errorf("device-camera-go: unrecognized write command")
}

// CameraRelease immediately returns control to the caller
func (n Client) CameraRelease(force bool) {
}

// CameraInit immediately returns control to the caller
func (n Client) CameraInit(edgexDevice models.Device, edgexProfile models.DeviceProfile, ipAddress string, username string, password string) {
}

// NewClient returns a new noop Client
func NewClient() client.Client {
	return Client{}
}
