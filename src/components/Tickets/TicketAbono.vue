<template>
  <v-layout class="ticket" wrap row>
    <v-flex xs12 class="hidden-print-only">
      <encabezado></encabezado>
      <v-btn small color="success" @click="volver()">
        <v-icon>arrow_back</v-icon>
        Volver
      </v-btn>
    </v-flex>
    <v-flex xs12>
      <br />
      <v-img class="text-xs-center" :src="require('@/assets/inicio/logo.png')">
      </v-img>
      <div class="text-xs-center">
        <p>
          <strong>Ticket de abono #{{ abono.IdAbono }}</strong>
          <br />
          <strong>Apartado #{{ abono.IdApartado }}</strong>
        </p>
        <p v-if="ajustes.Nombre">
          {{ ajustes.Nombre }}
        </p>
        <p v-if="ajustes.Direccion">
          {{ ajustes.Direccion }}
        </p>
        <p v-if="ajustes.Telefono">Teléfono: {{ ajustes.Telefono }}</p>
        <br />
        <p>
          {{ abono.Fecha | fechaExpresiva }}
        </p>
        <p>
          Lo atendió: <strong>{{ abono.Usuario.Nombre }}</strong>
        </p>
        <p>
          Cliente: <strong>{{ apartado.Cliente.Nombre }}</strong>
        </p>
        <div class="text-xs-center">
          <v-flex xs9 offset-xs1 class="text-xs-right con-borde-separador">
            <br />
          </v-flex>
        </div>
      </div>
      <v-layout wrap row>
        <template v-for="producto in apartado.Productos">
          <v-flex xs12 class="text-xs-left">{{ producto.Descripcion }}</v-flex>
          <v-flex xs12 class="text-xs-right con-borde-inferior"
            >{{ producto.Cantidad }} x {{ producto.PrecioVenta | currency }}
            =
            {{ (producto.Cantidad * producto.PrecioVenta) | currency }}
          </v-flex>
        </template>
      </v-layout>
      <div class="text-xs-center">
        <v-flex xs9 offset-xs1 class="text-xs-right con-borde-separador">
          <br />
        </v-flex>
      </div>
      <div class="text-xs-right">
        <p><strong>Total</strong> {{ apartado.Total | currency }}</p>
        <p>
          <strong>Restante anterior</strong>
          {{
            (apartado.Total -
              apartado.Abonado -
              apartado.Anticipo +
              abono.Monto)
              | currency
          }}
        </p>
        <p><strong>Cantidad abonada</strong> {{ abono.Monto | currency }}</p>
        <p><strong>Su pago</strong> {{ abono.Pago | currency }}</p>
        <p>
          <strong>Cambio</strong> {{ (abono.Pago - abono.Monto) | currency }}
        </p>
        <p>
          <strong>Restante actual</strong>
          {{
            (apartado.Total - apartado.Abonado - apartado.Anticipo) | currency
          }}
        </p>
        <br />
      </div>

      <div class="text-xs-center">
        <p v-if="ajustes.MensajePersonal">
          <strong>{{ ajustes.MensajePersonal }}</strong>
        </p>
        <Pie></Pie>
      </div>
    </v-flex>
    <v-btn
      :loading="cargando"
      class="hidden-print-only"
      @click="imprimir()"
      fixed
      dark
      fab
      bottom
      fill-height
      slot="activator"
      right
      color="green"
    >
      <v-icon>print</v-icon>
    </v-btn>
  </v-layout>
</template>

<script>
import { HTTP_AUTH } from "../../http-common";
import { EventBus } from "../../main";
import Encabezado from "./Encabezado";
import Pie from "./Pie";
import { TimeoutOcultarMenuTickets } from "../../constantes";

export default {
  name: "TicketAbono",
  components: { Pie, Encabezado },
  beforeRouteUpdate(detallesRuta) {
    this.obtenerDetallesDeAbono(
      detallesRuta.params.idAbono,
      detallesRuta.params.idApartado
    );
  },
  beforeMount() {
    EventBus.$emit("ponerTitulo", "Impresión de ticket");
    this.obtenerDetallesDeAbono(
      this.$route.params.idAbono,
      this.$route.params.idApartado
    );
  },
  data: () => ({
    cargando: false,
    apartado: {
      Usuario: {},
      Cliente: {},
    },
    abono: {
      Usuario: {},
    },
    ajustes: {},
  }),
  methods: {
    obtenerDetallesDeAbono(idAbono, idApartado) {
      if (!idApartado || !idAbono) return this.$router.go(-1);
      this.cargando = true;
      HTTP_AUTH.get("ajustes/empresa")
        .then((ajustes) => {
          this.ajustes = ajustes;
        })
        .then(() =>
          HTTP_AUTH.get(`apartado/${idApartado}`).then((apartado) => {
            this.apartado = apartado;
          })
        )
        .then(() =>
          HTTP_AUTH.get(`abono/${idAbono}/${idApartado}`)
            .then((abono) => {
              this.abono = abono;
            })
            .finally(() => (this.cargando = false))
        );
    },
    imprimir() {
      if (this.cargando) return;
      EventBus.$emit("ocultarMenu");
      // Esperar a que el menú esté oculto
      setTimeout(() => {
        let tituloOriginal = document.title;
        document.title = `Apartado #${this.apartado.Numero}`;
        window.print();
        document.title = tituloOriginal;
        EventBus.$emit("mostrarMenu");
      }, TimeoutOcultarMenuTickets);
    },
    volver() {
      this.$router.go(-1);
    },
  },
};
</script>
