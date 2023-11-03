<template>
  <div class="register">
    <h2>User Register</h2>
    <form @submit.prevent="registerUser">
      <div class="sign-field">
        <label>Name:</label>
        <input v-model="user.name" />
      </div>

      <div class="sign-field">
        <label>LoginID:</label>
        <input v-model="user.name" type="text">
      </div>

      <div class="sign-field">
        <label>Email:</label>
        <input v-model="user.email" type="email" />
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
        name: '',
        email: '',
        login_id: '',
        password: '',
      },
    };
  },
  methods: {
    registerUser() {
      // Handle user registration logic
      console.log('User registered:', this.user);
    },
  },
};
</script>

<style>
/* Add styles as needed */
.register {
  width: 300px;
}

.sign-field label {
  width: 100px;
}
</style>
