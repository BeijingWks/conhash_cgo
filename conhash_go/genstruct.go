package conhash

// #include "../cconhash/configure.h"
// #include "../cconhash/conhash_inter.h"
// #include "../cconhash/conhash.h"
// #include "../cconhash/util_rbtree.h"
import "C"

type conHashHandle C.struct_conhash_s
type conHashNode C.struct_node_s
type rbTreeNode C.struct_util_rbtree_node_s
type rbTree C.struct_util_rbtree_s

//go tool cgo -godefs genstruct.go
