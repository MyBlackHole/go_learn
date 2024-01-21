msgp -file=user.go

dlv test . -- -test.v -test.run TestMarshalUnmarshalPerson
