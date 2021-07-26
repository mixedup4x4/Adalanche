package main

import (
	"github.com/gofrs/uuid"
	"github.com/mattn/go-colorable"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Truly horrible, shoot me already

var (
	randguid, _    = uuid.NewV4()
	AttackerObject = &Object{
		DistinguishedName: "Attacker",
		Attributes: map[Attribute][]string{
			Name:       {"Attacker"},
			ObjectSid:  {string(AttackerSID)},
			ObjectGUID: {string(randguid.Bytes())},
		},
	}
	AllObjects              Objects
	SecurityDescriptorCache = make(map[uint32]*SecurityDescriptor)
	AllRights               = make(map[uuid.UUID]*Object) // Extented-Rights from Configuration - rightsGUID -> object
	AllSchemaClasses        = make(map[uuid.UUID]*Object) // schemaIdGUID -> object
	AllSchemaAttributes     = make(map[uuid.UUID]*Object) // attribute ...
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: colorable.NewColorableStdout()})
	AllObjects.Init("")
	AllObjects.Add(AttackerObject)
}
