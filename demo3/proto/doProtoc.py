import os

# buildSh = 'protoc -I . --grpc-gateway_out=. --go_out=plugins=grpc:. ./*.proto'
# -I 指定 proto 文件目录
# --grpc-gateway_out 输出 http 代码
# --go_out 生成 go 的目录
# ./*.proto 指定proto文件
# buildSh = 'protoc -I . --grpc-gateway_out=. --go-grpc_out=. --go_out=. ./*.proto'
buildSh = 'protoc --experimental_allow_proto3_optional -I .\
         --go_out=paths=source_relative:. \
         ./*.proto'

def is_dir(path):
    if os.path.isdir(path):
        print ("it's a directory")
        return True
    elif os.path.isfile(path):
        print ("it's a normal file")
        return False
    else:
        print ("it's a special file(socket,FIFO,device file)")
        return False
def build_unit(path):
    if is_dir(path):
        print('build_unit dir ：'+path)
        os.chdir(path)
        val = os.popen(buildSh)
        out = val.read()
        print(out)
    else:
        val = os.popen(buildSh)
        out = val.read()
        print(out)


def build_all():
    path = os.getcwd()
    print('-----')
    print(path)
    print('-----')
    proto_path = path
    doc = os.listdir('.')
    print('-----')
    print(doc)
    print('-----')
    for i in doc:
        print(i)
        son_path = proto_path + '/' + i
        print('work on ', son_path)
        build_unit(son_path)


# test_path = '/home/pines/code/ClearGrass/CommonComponents/protobuf/example'
# build_unit(test_path)

build_all()