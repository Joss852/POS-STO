<template>

  <div>
    <dialogo-cambiar-precio-venta
      @cerrar="dialogos.cambiarPrecio = false"
      :producto="productoParaCambiarPrecioDeVenta"
      :mostrar="dialogos.cambiarPrecio"
    ></dialogo-cambiar-precio-venta>
    <v-text-field
      @click:prepend="quiereBuscarProducto"
      label="Escanear código o escribir el número y presionar Enter"
      v-model="codigo"
      @keyup.enter="buscarYAgregarProductoSiExiste()"
      prepend-icon="search"
      solo
      clearable></v-text-field>
    <br/>
    <v-btn title="Categorias" icon class="mx-0" @click="cambiarPrecio(props.item.Numero)">
      <v-icon color="blue">delete</v-icon>
    </v-btn>
    <h1 class="title">Total de la venta: {{totalVenta | currency}}</h1>
    <br/>

    <div class="contenedor-lista-de-productos">
      <v-data-table :headers="headers" :items="listaDeProductos" hide-actions class="elevation-1">

        <template slot="items" slot-scope="props">
          <td>{{ props.item.Numero }}</td>
          <td>{{ props.item.CodigoBarras }}</td>
          <td>{{ props.item.Descripcion }}</td>
          <td>{{ props.item.Cantidad}}</td>
          <td>{{ props.item.PrecioVenta | currency}}</td>
          <td>{{ props.item.Total | currency}}</td>
          <td class="justify-center layout px-0">
            <v-btn title="Cambiar precio" icon class="mx-0" @click="cambiarPrecio(props.item.Numero)">
              <v-icon color="blue">edit</v-icon>
            </v-btn>
            <v-btn title="Aumentar uno" icon class="mx-0" @click="aumentarCantidad(props.item.Numero)">
              <v-icon color="green">add_shopping_cart</v-icon>
            </v-btn>
            <v-btn title="Restar uno" icon class="mx-0" @click="disminuirCantidad(props.item.Numero)">
              <v-icon color="orange">remove_shopping_cart</v-icon>
            </v-btn>
            <v-btn title="Quitar de la lista" icon class="mx-0" @click="quitarDeLaLista(props.item.Numero)">
              <v-icon color="red">delete</v-icon>
            </v-btn>
          </td>
        </template>
        <template slot="no-data">
          <v-alert :value="true" color="info">
            <div class="text-xs-center">
              <h1>Lista vacía</h1>
              <v-icon class="icono-grande">shopping_basket</v-icon>
              <p>Aquí aparecerá la lista de productos para la venta</p>
            </div>
          </v-alert>
        </template>
      </v-data-table>
    </div>
  </div>
</template>
<script>
  import {HTTP_AUTH} from "../../http-common";
  import DialogoCambiarPrecioVenta from './DialogoCambiarPrecioVenta'
  import Publicidad from "../Publicidad";

  export default {
    components: {Publicidad, DialogoCambiarPrecioVenta},
    beforeMount() {
      this.obtenerModoDeLecturaDeCodigosDeBarras();
    },
    computed: {
      totalVenta() {
        if (this.listaDeProductos.length <= 0) return 0;
        return this.listaDeProductos.reduce(
          (acumulador, siguiente) => ({
            Total: acumulador.Total + siguiente.Total
          }),
          {
            Total: 0
          }
        ).Total;
      }
    },
    data: () => ({
      urlBase: "producto",
      dialogos: {
        cambiarPrecio: false,
      },
      productoParaCambiarPrecioDeVenta: {},
      codigo: null,
      fab: false,
      listaDeProductos: [],
      headers: [
        {
          text: "#",
          align: "left",
          value: "Numero"
        },
        {
          text: "Código de barras",
          align: "left",
          value: "CodigoBarras"
        },
        {
          text: "Descripción",
          value: "Descripcion"
        },
        {
          text: "Cantidad",
          value: "Cantidad"
        },
        {
          text: "Precio",
          value: "PrecioVenta"
        },
        {
          text: "Total",
          value: "Total"
        },
        {
          text: "Opciones",
          sortable: false
        }
      ]
    }),
    methods: {
      cambiarPrecio(numero) {
        this.productoParaCambiarPrecioDeVenta = this.listaDeProductos.find(producto => producto.Numero === numero);
        this.dialogos.cambiarPrecio = true;

      },
      obtenerModoDeLecturaDeCodigosDeBarras() {
        HTTP_AUTH.get("ajustes/otros").then(ajustes => {
          let {ModoLecturaProductos} = ajustes;
          if (ModoLecturaProductos === "codigo") this.urlBase = "producto/codigo/barras";
        });
      },
      quiereBuscarProducto() {
        this.$emit("buscar-producto");
      },
      cancelarVenta() {
        this.listaDeProductos = [];
      },
      aumentarCantidad(numero) {
        let producto = this.listaDeProductos.find(producto => producto.Numero === numero);
        if (!producto) return;
        if (producto.Existencia > producto.Cantidad) {
          this.$set(producto, "Cantidad", producto.Cantidad + 1);
          this.$set(producto, "Total", producto.PrecioVenta * producto.Cantidad);
        }
      },
      disminuirCantidad(numero) {
        let producto = this.listaDeProductos.find(producto => producto.Numero === numero);
        if (!producto) return;
        if (producto.Cantidad > 1) {
          this.$set(producto, "Cantidad", producto.Cantidad - 1);
          this.$set(producto, "Total", producto.PrecioVenta * producto.Cantidad);
        }
      },
      quitarDeLaLista(numero) {
        let indice = this.listaDeProductos.findIndex(producto => producto.Numero === numero);
        if (indice === -1) return;
        this.listaDeProductos.splice(indice, 1);
      },
      agregarOModificarExistenciaDeProductoEnLista(producto) {
        const indice = this.listaDeProductos.findIndex(
          productoExistente => productoExistente.Numero === producto.Numero
        );
        if (indice === -1) {
          this.listaDeProductos.push({
            Numero: producto.Numero,
            Descripcion: producto.Descripcion,
            Cantidad: 1,
            CodigoBarras: producto.CodigoBarras,
            Existencia: producto.Existencia,
            PrecioVenta: producto.PrecioVenta,
            PrecioVentaOriginal: producto.PrecioVenta,
            PrecioCompra: producto.PrecioCompra,
            Total: producto.PrecioVenta
          });
        } else {
          let productoExistente = this.listaDeProductos[indice];
          if (productoExistente.Cantidad < productoExistente.Existencia) {
            productoExistente.Cantidad++;
            productoExistente.Total =
              productoExistente.Cantidad * productoExistente.PrecioVenta;
          }
        }
      },
      buscarYAgregarProductoSiExiste() {
        if (this.codigo) {
          HTTP_AUTH.get(`${this.urlBase}/${this.codigo}`).then(producto => {
            if (null !== producto) {
              this.agregarOModificarExistenciaDeProductoEnLista(producto)
            } else {
              this.$emit("producto-no-existente");
            }
            this.codigo = null;
          });
        }
      }
    }
  };
</script>
