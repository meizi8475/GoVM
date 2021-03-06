package reserved

import (
	"GoVM/chapter5-instructions/base"
	"GoVM/chapter4-rtdt"
	"GoVM/native"
	//如果不显式使用lang包中的变量，只是让他执行init()方法，需要前面加下划线
	_ "GoVM/native/java/lang"
	_ "GoVM/native/sun/misc"
)

/**
	即0xFE指令
 */
type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKE_NATIVE) Execute(frame *chapter4_rtdt.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "." + methodName + "--" + methodDescriptor
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	//这相当于调用找到的native method
	nativeMethod(frame)
}