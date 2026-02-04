package nbt

import "fmt"

var TagName = map[tagTypeByte]string{
	BTagCompound:  "TAG_Compound",
	BTagEnd:       "TAG_End",
	BTagByte:      "TAG_Byte",
	BTagShort:     "TAG_Short",
	BTagInt:       "TAG_Int",
	BTagLong:      "TAG_Long",
	BTagFloat:     "TAG_Float",
	BTagDouble:    "TAG_Double",
	BTagByteArray: "TAG_Byte_Array",
	BTagString:    "TAG_String",
	BTagList:      "TAG_List",
	BTagIntArray:  "TAG_Int_Array",
	BTagLongArray: "TAG_Long_Array",
}

func PrintTag(tag NBTTag) {
	fmt.Println("Tag Type:", TagName[tag.Type()])
	fmt.Println("Tag Name:", tag.Name())
	fmt.Println("Tag Data Length:", tag.DataLength())
	fmt.Println("Z index:", tag.ZIndex())
	switch tag.Type() {
	case BTagByte:
		fmt.Println("Tag Value:", tag.(*TagByte).Value)
	case BTagShort:
		fmt.Println("Tag Value:", tag.(*TagShort).Value)
	case BTagInt:
		fmt.Println("Tag Value:", tag.(*TagInt).Value)
	case BTagLong:
		fmt.Println("Tag Value:", tag.(*TagLong).Value)
	case BTagFloat:
		fmt.Println("Tag Value:", tag.(*TagFloat).Value)
	case BTagDouble:
		fmt.Println("Tag Value:", tag.(*TagDouble).Value)
	case BTagString:
		fmt.Println("Tag Value:", tag.(*TagString).Value)
	case BTagCompound:
		{
			for _, childTag := range tag.(*TagCompound).Value {
				PrintTag(childTag)
			}
		}
	default:
		fmt.Println("Tag Value: [Not Implemented]")
	}
	fmt.Println()
}
