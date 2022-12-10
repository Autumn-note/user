

function traversalFolder() {
    for file in `ls $1`
    do
        if [ -d $1"/"$file ]
        then
            traversalFolder $1"/"$file
        else
            if [ "${file##*.}"x = "proto"x ]
            then
                echo $1"/"$file
                echo "current file is $file"
                protoc   --go-grpc_out=require_unimplemented_servers=false,paths=source_relative:. --go_out=paths=source_relative:. $1"/"$file
            fi
        fi
    done
}


traversalFolder $1