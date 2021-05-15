package main

//https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-02-pb-intro.html
//4.2.2 定制化代码生成插件,本次例子定制化netrpc 插件,根据/xx/generator interface来定制化
//每次更新这个定制化的文件都必须更新二进制文件(go build -o protoc-gen-go-netrpc main.go)
import (
	"bytes"
	"html/template"
	"log"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

//要使用自定义插件必须先注册插件
func init() {
	generator.RegisterPlugin(new(netrpcPlugin))
}

type netrpcPlugin struct {
	*generator.Generator
}

//插件名
func (p *netrpcPlugin) Name() string {
	return "netrpc"
}

//通过g参数对插件初始化
func (p *netrpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

//生成主体包代码
func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

//生成导入包代码
func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

//以下为自定义插件netrpc 生成文件的内容
func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P(`import "net/rpc"`)
}
func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	spec := p.buildServiceSpec(svc)

	var buf bytes.Buffer
	t := template.Must(template.New("").Parse(tmplService))
	err := t.Execute(&buf, spec)
	if err != nil {
		log.Fatal(err)
	}
	p.P(buf.String())
}

//要在自定义的genServiceCode中为每个服务生成相关的代码,定义一个ServiceSpec用于描述服务的元信息
type ServiceSpec struct {
	ServiceName string
	MethodList  []ServiceMethodSpec
}
type ServiceMethodSpec struct {
	MethodName     string
	InputTypeName  string
	OutputTypeName string
}

//新建一个buildServiceSpec 来解析每个服务的ServiceSpec元信息
func (p *netrpcPlugin) buildServiceSpec(svc *descriptor.ServiceDescriptorProto) *ServiceSpec {
	//svc.GetName() 可以获取protobuf文件中定义的服务(service xxx {})的名字
	spec := &ServiceSpec{
		ServiceName: generator.CamelCase(svc.GetName()),
	}

	for _, m := range svc.Method {
		spec.MethodList = append(spec.MethodList, ServiceMethodSpec{
			MethodName:     generator.CamelCase(m.GetName()),
			InputTypeName:  p.TypeName(p.ObjectNamed(m.GetInputType())),
			OutputTypeName: p.TypeName(p.ObjectNamed(m.GetOutputType())),
		})
	}

	return spec
}

var tmplService = `
{{$root := .}}

type {{.ServiceName}}Interface interface {
    {{- range $_, $m := .MethodList}}
    {{$m.MethodName}}(*{{$m.InputTypeName}}, *{{$m.OutputTypeName}}) error
    {{- end}}
}

func Register{{.ServiceName}}(
    srv *rpc.Server, x {{.ServiceName}}Interface,
) error {
    if err := srv.RegisterName("{{.ServiceName}}", x); err != nil {
        return err
    }
    return nil
}

type {{.ServiceName}}Client struct {
    *rpc.Client
}

var _ {{.ServiceName}}Interface = (*{{.ServiceName}}Client)(nil)

func Dial{{.ServiceName}}(network, address string) (
    *{{.ServiceName}}Client, error,
) {
    c, err := rpc.Dial(network, address)
    if err != nil {
        return nil, err
    }
    return &{{.ServiceName}}Client{Client: c}, nil
}

{{range $_, $m := .MethodList}}
func (p *{{$root.ServiceName}}Client) {{$m.MethodName}}(
    in *{{$m.InputTypeName}}, out *{{$m.OutputTypeName}},
) error {
    return p.Client.Call("{{$root.ServiceName}}.{{$m.MethodName}}", in, out)
}
{{end}}
`
