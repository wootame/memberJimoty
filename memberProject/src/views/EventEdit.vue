<template>
  <div class="event-form main-stage">
    <h2>{{ isEdit ? 'イベント編集' : '新規イベント作成' }}</h2>
    <form @submit="submitEvent">
      <div class="form-group">
        <label for="title">タイトル</label>
        <input type="text" id="title" v-model="event.title" />
      </div>
      <div class="form-group">
        <label for="category">カテゴリー</label>
        <select id="category" v-model="event.selectedCategory">
          <option v-for="cat in category" :value="cat" :key="cat">{{ cat }}</option>
        </select>
      </div>
      <div class="form-group">
        <label for="atmosphere">雰囲気</label>
        <input type="text" id="atmosphere" v-model="event.atmosphere" />
      </div>
      <div class="form-group">
        <label for="tags">タグ</label>
        <input type="text" id="tags" v-model="newTag"/>
        <button @click.prevent="addTag">追加</button>
        <div class="event-tags">
          <span v-for="(tag, index) in event.tags" :key="index">{{ tag }}</span>
        </div>
      </div>
      <div class="form-group">
        <label for="description">説明</label>
        <textarea id="description" v-model="event.description"></textarea>
      </div>
      <input :value="isEdit ? '更新' : '作成'" type="submit">
    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      event: {
        id: null,
        title: 'test',
        selectedCategory: "イベント",
        atmosphere: '',
        tags: [
          'テスト',
          'テスト'
        ],
        description: 'test',
      },
      category: [
        'イベント',
        '飲み会',
        'スポーツ',
      ],
      newTag: '',
      isEdit: false,
    };
  },
  methods: {
    submitEvent() {
      // Handle form submission, create or update event
      if (this.isEdit) {
        // Update event logic here
      } else {
        // Create event logic here
      }
      // After creating/updating the event, you can redirect to the event list or do other actions.
    },
    addTag() {
      if (this.newTag.trim() !== '') {
        this.event.tags.push(this.newTag);
        this.newTag = ''; // 入力フィールドをクリア
      }
    },
  },
  created() {
    // Check if this is an edit page or create page and initialize the event data accordingly.
    // If it's an edit page, set this.isEdit to true and fetch the event data to populate the form.
    // If it's a create page, this.isEdit remains false.
  },
};
</script>

<style scoped>
  .event-form {
    width: 80%;
    border: 1px solid #888;
    border-radius: 5px;
    padding: 0 20px 20px;
    margin: 20px auto;
  }

  .form-group {
    margin: 20px 0;
  }

  .form-group label {
    font-size: 22px;
    display: block;
  }

  .form-group input[type="text"],
  .form-group textarea,
  select {
    width: 80%;
    font-size: 20px;
    border: solid 1px #888;
    border-radius: 5px;
    margin-left: 8%;
  }

  .event-tags {
    margin-top: 10px;
    margin-left: 8%;
  }

  .event-tags span {
    background-color: #888;
    font-size: 18px;
    color: #fff;
    border-radius: 5px;
    padding: 5px;
  }

  .event-tags span + span {
    margin-left: 10px;
  }

  .form-group input[type="text"],
  select {
    height: 40px;
  }

  #tags {
    width: 70%;
  }

  .form-group button {
    width: 6%;
    text-align: center;
    margin-left: 2%;
  }

  .form-group textarea {
    min-height: 200px;
  }

  input[type="submit"] {
    background-color: #888;
    width: 200px;
    height: 40px;
    text-align: center;
    color: #fff;
    border-radius: 5px;
    margin: 0 auto;
    display: block;
  }
</style>