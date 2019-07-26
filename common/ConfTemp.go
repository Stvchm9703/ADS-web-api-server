package common

type ConfigTemp struct {
	Server struct {
		IP               string `toml:"ip" json:"ip" yaml:"ip"`
		Port             int    `toml:"port" json:"port" yaml:"port"`
		RootFilePath     string `toml:"root_path" json:"root_path" yaml:"root_path"`
		MainPath         string `toml:"main_path" json:"main_path" yaml:"main_path"`
		StaticFilepath   string `toml:"static_filepath" json:"static_filepath" yaml:"static_filepath"`
		StaticOutpath    string `toml:"static_outpath" json:"static_outpath" yaml:"static_outpath"`
		TemplateFilepath string `toml:"template_filepath" json:"template_filepath" yaml:"template_filepath"`
		TemplateOutpath  string `toml:"template_outpath" json:"template_outpath" yaml:"template_outpath"`
		APITablePath     string `toml:"api_table_filepath" json:"api_table_filepath" yaml:"api_table_filepath"`
		APIOutpath       string `toml:"api_outpath" json:"api_outpath" yaml:"api_outpath"`
	} `toml:"prod_server" json:"prod_server" yaml:"prod_server"`

	Database struct {
		Connector string `toml:"connector" json:"connector" yaml:"connector"`
		Host      string `toml:"host" json:"host" yaml:"host"`
		Port      int    `toml:"port" json:"port" yaml:"port"`
		Username  string `toml:"username" json:"username" yaml:"username"`
		Password  string `toml:"password" json:"password" yaml:"password"`
		Database  string `toml:"database" json:"database" yaml:"database"`
		Filepath  string `toml:"filepath" json:"filepath" yaml:"filepath"`
	} `toml:"database" json:"database" yaml:"database"`

	DatabaseFallback struct {
		// .csv / .db
		Schema   string `toml:"schema" json:"schema" yaml:"schema"`
		Filepath string `toml:"filepath" json:"filepath" yaml:"filepath"`
	} `toml:"db_fallback" json:"db_fallback" yaml:"db_fallback"`
}
