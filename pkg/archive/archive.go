package archive

type FileType int32
const (
	Folder string = "folder"
)

func NewArchiveService() *archive{
	return &archive {
		DefaultFolder: []string{},
	}
}

type archive struct {
	DefaultFolder []string
}

func (a *archive) Init() {

}

//创建默认文件夹
func (a *archive) createDefaultFolder() {

}

