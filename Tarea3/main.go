package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
)

type nodo struct {
	nombre, apellido, apodo, favoritos string
	Siguiente, Anterior *nodo
}

type lista struct {
	cabeza *nodo
	cola   *nodo
}

func (this *lista) Insertar(nuevo *nodo) {
	if this.cabeza == nil {
		this.cabeza = nuevo
		this.cola = nuevo
	} else {
		this.cola.Siguiente = nuevo
		nuevo.Anterior = this.cola
		this.cola = nuevo
	}
}

func graficar(l lista){
	arch, _ := os.Create("archivo.dot")
	//Encabezado
	_, _ = arch.WriteString("digraph G{" + "\n")
	_, _ = arch.WriteString(`compound=true;` + "\n")
	_, _ = arch.WriteString(`subgraph cluster1{style=invis;` + "\n")
	_, _ = arch.WriteString(`edge[dir=both]` + "\n")

	//Creacion Tiendas
	aux := l.cabeza
	cant := 1
	for aux != nil{
		_, _ = arch.WriteString("nodo" + strconv.Itoa(cant) + `[shape=record,color=green,label="{{` + aux.nombre + " | " + aux.apellido + " } |" + aux.apodo + " | " + aux.favoritos + `}"];` + "\n")
		cant++
		aux = aux.Siguiente
	}

	//Cantidad de Tiendas
	contador := 0
	aux = l.cabeza
	for aux != nil{
		contador++
		aux = aux.Siguiente
	}

	//Conexiones Tiendas
	caux := cant-contador
	canta := caux
	aux = l.cabeza
	conaux := 1
	for aux != nil{
		canta++
		if contador != 1 && conaux < contador{
			_, _ = arch.WriteString("nodo" + strconv.Itoa(caux) + " -> nodo" + strconv.Itoa(canta)  + "; \n")
			_, _ = arch.WriteString("nodo" + strconv.Itoa(canta) + " -> nodo" + strconv.Itoa(caux)  + "; \n")
		}
		conaux++
		caux++
		aux = aux.Siguiente
	}
	_, _ = arch.WriteString("}" + "\n")
	_, _ = arch.WriteString("}" + "\n")

	path, _ := exec.LookPath("dot")
	cmd, _ := exec.Command(path, "-Tpng", "./archivo.dot").Output()
	mode := 0777
	_ = ioutil.WriteFile("outfile.png", cmd, os.FileMode(mode))
}

func main() {
	li := lista{nil, nil}
	a := nodo{"Marvin", "Martinez", "Marvin25ronal", "Jugar apex", nil, nil}
	b := nodo{"Yaiza", "Pineda", "Bambi", "Patinar", nil, nil}
	c := nodo{"Jonathan", "Lopez", "Pancho", "Comer", nil, nil}
	d := nodo{"usuario1", "bla", "bla", "Jugar apex", nil, nil}
	e := nodo{"usuario2", "bla", "bla", "Jugar apex", nil, nil}
	f := nodo{"usuario3", "sale edd", "vamos con todo", "100 en la fase 1", nil, nil}
	li.Insertar(&a)
	li.Insertar(&b)
	li.Insertar(&c)
	li.Insertar(&d)
	li.Insertar(&e)
	li.Insertar(&f)
	graficar(li)
}
