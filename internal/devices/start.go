// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package devices

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Start() *cobra.Command {
	var deviceID string
	// startdeviceCmd represents the startdevice command
	var startDeviceCmd = &cobra.Command{
		Use:   `start -i <device_id>`,
		Short: "Starts a device.",
		Long:  "Starts or powers on a device that is currently stopped or powered off.",
		Example: `  # Starts the specified device:
  metal device start --id 26a9da5f-a0db-41f6-8467-827e144e59a7`,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			_, err := c.Service.PowerOn(deviceID)

			if err != nil {
				return errors.Wrap(err, "Could not start Device")
			}

			fmt.Println("Device", deviceID, "successfully started.")
			return nil
		},
	}

	startDeviceCmd.Flags().StringVarP(&deviceID, "id", "i", "", "The UUID of the device.")
	_ = startDeviceCmd.MarkFlagRequired("id")
	return startDeviceCmd
}
