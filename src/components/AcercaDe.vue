<template>
  <v-layout wrap row>
    <v-flex xs12>
      <h3 class="display-2">Acerca de</h3>
      <span class="subheading"
        >POS - STO. Sistema de ventas web. Seguro, rápido, confiable y
        amigable.</span
      >
      <div class="title mb-3">
        Made with ❤️ by
        <a href="https://github.com/Joss852" target="_blank">Josseph M</a>
      </div>
      <v-divider class="my-3"></v-divider>
    </v-flex>
  </v-layout>
</template>
<script>
import { EventBus } from "../main";
import { EnlaceCompartir } from "../constantes";

export default {
  beforeMount() {
    EventBus.$emit("ponerTitulo", "Acerca de");
  },
  data: () => ({
    pagina: EnlaceCompartir,
  }),
};
</script>
