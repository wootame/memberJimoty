<template>
  <header-component></header-component>
  <router-view />
</template>

<script>
import EventListComponent from './views/EventList.vue';
import HeaderComponent from './views/template/Header.vue';

export default {
  name: 'App',
  components: {
    HeaderComponent,
    EventListComponent,
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  color: #2c3e50;
}

</style>
