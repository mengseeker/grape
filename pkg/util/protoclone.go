package util

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func AnypbClone(src proto.Message, dst proto.Message) error {
	// defer func() {
	// 	println("=========================")
	// 	println(reflect.DeepEqual(src, dst))
	// 	fmt.Printf("%v\n", src)
	// 	fmt.Printf("%v\n", dst)
	// }()
	raw, err := anypb.New(src)
	if err != nil {
		return err
	}
	return raw.UnmarshalTo(dst)
}
