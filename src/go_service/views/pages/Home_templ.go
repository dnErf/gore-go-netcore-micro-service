// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "gore/views/layouts"

func Home() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"p-4\"><div><h1 class=\"\">micro service</h1><hr><div class=\"my-2 flex flex-row gap-4\"><button id=\"testService\" class=\"border p-2\">Test Service</button> <button class=\"border p-2\">Test Log</button></div></div><div class=\"flex w-full\"><div class=\"w-1/2\"><h2>data sent</h2><pre id=\"sent\" class=\"border px-4\"></pre></div><div class=\"w-0 h-4\"></div><div class=\"w-1/2\"><h2>data received</h2><pre id=\"received\" class=\"border px-4\"></pre></div></div></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.Base().Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<script>\r\n        let sent = document.getElementById(\"sent\")\r\n        let received = document.getElementById(\"received\")\r\n\r\n        let payloadTest = {\r\n            action: \"test\",\r\n            data: {\r\n                message: \"some kind of test\"\r\n            }\r\n        }\r\n        let payStubTest = {\r\n            action: \"test\",\r\n            data: {\r\n                message: \"some kind of test\"\r\n            }\r\n        }\r\n\r\n        sent.innerHTML = \"...\"\r\n        received.innerHTML = \"...\"\r\n\r\n        let testService = document.getElementById(\"testService\")\r\n        testService.addEventListener(\"click\", function() {\r\n            console.log(\"clicked test service\")\r\n            sent.innerHTML = JSON.stringify(payloadTest, undefined, 4)\r\n\r\n            fetch(\"/test-service\", {\r\n                method: \"POST\",\r\n                body: JSON.stringify(payloadTest)\r\n            })\r\n            .then((response) => response.json())\r\n            .then((data) => {\r\n                received.innerHTML = JSON.stringify(data, undefined, 4)\r\n            })\r\n            \r\n            console.log(\"===\")\r\n        })\r\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
