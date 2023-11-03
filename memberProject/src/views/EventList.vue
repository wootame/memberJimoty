<template>
  <div id="eventList" class="main-stage">
    <h2>イベント一覧</h2>
    <div>
      <div v-for="event in events" :key="event.id" class="event">
        <router-link :to="`/events/${event.id}`">
          <h3>タイトル: {{ event.title }}</h3>
        </router-link>
        <hr>
        <p>カテゴリー: {{ event.category }}</p>
        <p>タグ: 
          <span v-for="tag in event.tags" :key="tag">{{ tag }}, </span>
        </p>
        <p>説明: {{ event.description }}</p>
        <p>コメント数: {{ event.commentCount }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import HeaderComponent from './template/Header.vue';
export default {
  components: {
    HeaderComponent,
  },
  data() {
    return {
      events: [
        { 
          id: 1, 
          title: 'イベント 1' ,
          category: '飲み会',
          atmosphere: 'しっぽり',
          tags: ['飲み会', 'しっぽり', '友達づくり'],
          description: 'こんなイベントです！',
          commentCount: 15,
        },
        { 
          id: 2, 
          title: 'イベント 2',
          category: 'スポーツ',
          atmosphere: 'わいわい',
          tags: ['スポーツ', '健康意識', 'アクティブ'],
          description: 'こんなイベントです', 
          commentCount: 15,
        },
        { 
          id: 3, 
          title: 'イベント 1' ,
          category: '飲み会',
          atmosphere: 'ワイワイ',
          tags: ['飲み会', '楽しく', '大勢'],
          description: 'こんなイベントです！',
          commentCount: 15,
        },
      ],
    };
  },
};
</script>

<style>
/* Add styles as needed */
</style>
