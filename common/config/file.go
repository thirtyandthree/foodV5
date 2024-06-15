package config

type File struct {
	// 上传图片路径地址
	UploadDir string `yaml:"upload_dir"`
	// 自大上传空间
	UploadMaxSize int64 `yaml:"upload_max_size"`
	// 允许上传的资源类型后缀
	UploadAllowFiles []string `yaml:"upload_allow_files"`
}
