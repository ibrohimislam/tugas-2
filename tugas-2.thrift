namespace go tugas

struct File {
    1: required string filename,
    2: required string size,
    3: required string mode,
    4: required string modifiedTime,
    5: required string createdTime,
    6: required bool isDir
}

service FileSvc {
    list<File> Dir(1:string path)
    void CreateDir(1:string path, 2:string name)
    list<byte> GetContent(1:string path, 2:string name)
    string PutContentOpen(1:string path, 2:string name, 3:string hash)
    i32 PutContentPartial(1:string handler, 2:list<byte> content)
    i32 PutContentClose(1:string handler)
}