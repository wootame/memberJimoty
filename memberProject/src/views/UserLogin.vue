<template>
  <div>
    <h2>User Login</h2>
    <form @submit.prevent="loginUser">
      <div class="sign-field">
        <label>UserName:</label>
        <input v-model="user.login_id" type="text" />
      </div>

      <div class="sign-field">
        <label>Password:</label>
        <input v-model="user.password" type="password" />
      </div>

      <input type="submit">
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      user: {
        login_id: '',
        password: '',
      },
    };
  },
  methods: {
    loginUser() {
      console.log('User logged in:', this.user);
    },
  },
};
</script>

<style>
</style>
