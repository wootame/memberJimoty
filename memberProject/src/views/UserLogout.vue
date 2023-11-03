<template>
  <div>
    <h2>Logout</h2>
    <p>Are you sure you want to log out?</p>
    <button @click="logout">Logout</button>
  </div>
</template>

<script>
export default {
  methods: {
    logout() {
      // Add logic for user logout (e.g., clearing authentication state, redirecting, etc.)
      console.log('User logged out');
      // Example: Redirect to the login page after logout
      this.$router.push('/login');
    },
  },
};
</script>

<style>
/* Add styles as needed */
</style>
