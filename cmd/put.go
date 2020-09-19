/*
Copyright © 2020 Rafał Korkosz <korkosz.rafal@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/rkorkosz/rlog/internal/app/rlog"
	"github.com/rkorkosz/rlog/internal/pkg/editor"
	"github.com/rkorkosz/rlog/internal/pkg/slug"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Put entry",
	RunE: func(cmd *cobra.Command, args []string) error {
		storage, err := rlog.NewBoltStorage(viper.GetString("dbpath"), "entries")
		if err != nil {
			return err
		}
		input, err := editor.CaptureFromEditor()
		if err != nil {
			return err
		}
		spl := strings.SplitN(string(input), "\n", 2)
		e := rlog.Entry{
			Title: spl[0],
			Slug:  slug.Slug(spl[0]),
			Text:  strings.TrimSpace(spl[1]),
		}
		err = storage.Put(e)
		if err != nil {
			return err
		}
		return json.NewEncoder(os.Stdout).Encode(&e)
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}