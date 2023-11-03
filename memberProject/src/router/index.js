// src/router/index.js
import { createRouter, createWebHistory } from 'vue-router';

// ページコンポーネントのインポート
import EventList from '../views/EventList.vue';
import EventDetail from '../views/EventDetail.vue';
import EventEdit from '../views/EventEdit.vue';
import UserLogin from '../views/UserLogin.vue';
import UserLogout from '../views/UserLogout.vue';
import UserRegister from '../views/UserRegister.vue';
import UserProfile from '../views/UserProfile.vue';
import UserProfileEdit from '../views/UserProfileEdit.vue';
import ChatList from '../views/ChatList.vue';
import Chat from '../views/Chat.vue';

const routes = [
  // ログイン周り
  { path: '/login', component: UserLogin },
  { path: '/logout', component: UserLogout },
  { path: '/register', component: UserRegister },
  // ログイン後
  { path: '/', component: EventList },
  { path: '/events', component: EventList, meta: { requiresAuth: true } },
  { path: '/events/details/:id', component: EventDetail },
  { path: '/events/edit', component: EventEdit },
  { path: '/events/edit/:id', component: EventEdit },
  { path: '/user/profile/:id', component: UserProfile },
  { path: '/user/edit/:id', component: UserProfileEdit },
  { path: '/chat', component: ChatList },
  { path: '/chat/:id', component: Chat },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  if (to.meta.requiresAuth) {
    authenticateUser()
      .then(() => {
        next();
      })
      .catch(() => {
        next('/login');
      })
  } else {
    next();
  }
});

function authenticateUser() {
  return new Promise((resolve, reject) => {

  });
}

export default router;
