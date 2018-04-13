package yaml

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"../../base"
	"gopkg.in/yaml.v2"
)

type Env struct {
	Selected string
	Excluded []string
}

func (env Env) ShouldIncludeFile(folder, file string) bool {
	//fmt.Printf("len excluded & selected is %d %d\n", len(env.Excluded), len(env.Selected))
	if len(env.Excluded) == 0 || len(env.Selected) == 0 {
		return true
	}
	//fmt.Printf("[debug] env should include %s %s\n", folder, file)

	if strings.HasSuffix(folder, env.Selected) {
		return true
	}

	for _, item := range env.Excluded {
		if strings.HasSuffix(folder, item) {
			return false
		}
	}

	return true
}

type Src struct {
	Name   string
	Suffix string
}

func (src Src) GetName() string {
	return src.Name
}

func (src Src) GetSuffix() string {
	return src.Suffix
}

type NsProcessor struct {
	OldNS string `yaml:"old-ns"`
	NewNS string `yaml:"new-ns"`
}
type Dest struct {
	Tpe             string      `yaml:"type"`
	Name            string
	NsProcessor     NsProcessor `yaml:"ns-processor"`
	RemoveCopyright bool        `yaml:"remove-copyright"`
	FileLineLimit   int         `yaml:"file-line-limit"`
	RemoveComment   bool        `yaml:"remove-comment"`
}

func (dest Dest) GetType() string {
	return dest.Tpe
}

func (dest Dest) GetName() string {
	return dest.Name
}

func (dest Dest) GetOldNS() string {
	return dest.NsProcessor.OldNS
}

func (dest Dest) GetNewNS() string {
	return dest.NsProcessor.NewNS
}

func (dest Dest) IsRemoveCopyright() bool {
	return dest.RemoveCopyright
}

func (dest Dest) GetFileLineLimit() int {
	return dest.FileLineLimit
}

func (dest Dest) ShouldRemoveComment() bool {
	return dest.RemoveComment
}

type Conf struct {
	Env  Env
	Src  Src
	Dest Dest
}

func (conf *Conf) GetEnv() base.EnvInterface {
	return conf.Env
}

func (conf *Conf) GetSrc() base.SrcInterface {
	return conf.Src
}

func (conf *Conf) GetDest() base.DestInterface {
	return conf.Dest
}

type ConfigOption struct {
	name string
}

func (cfg *ConfigOption) Set(s string) error {
	cfg.name = s
	return nil
}

func (cfg *ConfigOption) String() string {
	return fmt.Sprintf("-config %s", cfg.name)
}

func ConfigOptionFlag() *ConfigOption {
	cfg := &ConfigOption{}
	flag.CommandLine.Var(cfg, "config", "指定配置文件")
	return cfg
}

func NewContext(cfg *ConfigOption) (*Conf, error) {
	if len(cfg.name) == 0 {
		flag.Usage()
		return nil, fmt.Errorf("no config provided")
	}

	source, err := ioutil.ReadFile(cfg.name)
	if err != nil {
		return nil, err
	}

	conf := Conf{}

	err = yaml.Unmarshal(source, &conf)

	return &conf, nil
}

func ShowExample() {
	conf := Conf{
		Env: Env{
			Selected: "dev",
			Excluded: []string{"dev", "uat", "prod", "mars", "explorer"},
		},
		Src: Src{
			Name: "./",
		},
		Dest: Dest{
			Name: "/tmp/maycur",
			Tpe:  "batch",
			NsProcessor: NsProcessor{
				OldNS: "maycur",
				NewNS: "dev",
			},
			RemoveCopyright: true,
			FileLineLimit:   4000,
		},
	}

	d, err := yaml.Marshal(&conf)
	if err != nil {
		panic(err)
	}

	s := string(d)
	fmt.Println(s)
}
