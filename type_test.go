package luar

import (
	"testing"

	"github.com/zyedidia/gopher-lua"
)

func Test_type_slice(t *testing.T) {
	L := lua.NewState()
	defer L.Close()

	type ints []int

	L.SetGlobal("newInts", NewType(L, ints{}))

	testReturn(t, L, `ints = newInts(1); return #ints`, "1")
	testReturn(t, L, `ints = newInts(0, 10); return #ints`, "0")
}

func Test_type(t *testing.T) {
	L := lua.NewState()
	defer L.Close()

	tim := &StructTestPerson{
		Name: "Tim",
	}

	L.SetGlobal("user1", New(L, tim))
	L.SetGlobal("Person", NewType(L, StructTestPerson{}))
	L.SetGlobal("People", NewType(L, map[string]*StructTestPerson{}))

	testReturn(t, L, `user2 = Person(); user2.Name = "John"; user2.Friend = user1`)
	testReturn(t, L, `return user2.Name`, "John")
	testReturn(t, L, `return user2.Friend.Name`, "Tim")
	testReturn(t, L, `everyone = People(); everyone["tim"] = user1; everyone["john"] = user2`)

	everyone := L.GetGlobal("everyone").(*lua.LUserData).Value.(map[string]*StructTestPerson)
	if len(everyone) != 2 {
		t.Fatalf("expecting len(everyone) = 2, got %d", len(everyone))
	}
}

func Test_type_metatable(t *testing.T) {
	L := lua.NewState()
	defer L.Close()

	L.SetGlobal("newInt", NewType(L, int(0)))

	testReturn(t, L, `return getmetatable(newInt) == "ktluar"`, "true")
}


