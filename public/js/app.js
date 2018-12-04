window.Event = new Vue();
Vue.component('pembeli-modal', {
    data: function() {
        return {
            person: '',
            active: false
        }
    },
    methods: {
        show() {
            this.active = true;
        },
        hide() {
            this.active = false;
        }
    },
    created() {
        Event.$on('togglePembeliModal', (row) => {
          console.log(row);
          this.person = row;
          this.show();
        });
    },
    template: `<div class="modal" :class="{ 'is-active': active}">
      <div class="modal-background" v-on:click='hide()'></div>
      <div class="modal-content">
        <div class="card">
            <div class="card-content">
                <h1>Nama: {{person.nama}}</h1>
                <h1>Email: {{person.email}}</h1>
                <h1>Alamat: {{person.alamat}}</h1>
                <h1>Telepon: {{person.telepon}}</h1>
            </div>
        </div>
      </div>
      <button class="modal-close is-large" aria-label="close"></button>
    </div>`
})

Vue.component ('update-status', {
    data: function() {
        return {
            id: '',
            status: '',
            active: false,
            options: ['diterima', 'diproses', 'dikirimkan']
        }
    },
    methods: {
        show() {
            this.active = true;
        },
        update() {
            axios.post('/update-status', {
                id: this.id,
                status: this.status
            })
              .then(results => {this.active = false; Event.$emit('update');})
              .catch(err => console.log(err))
        },
        hide() {
            this.status = ''
            this.id = ''
            this.active = false;
        }
    },
    created() {
        Event.$on('toggleUpdateModal', (row) => {
          console.log(row);
          this.id = row;
          this.show();
        });
    },
    template: `<div class="modal" :class="{ 'is-active': active}">
      <div class="modal-background" v-on:click='hide()'></div>
      <div class="modal-content">
        <div class="card">
            <div class="card-content">
                <input class="input" type="text" placeholder="Text input" v-model="status">
                <a class="button" v-on:click='update()'>Update</a>
            </div>
        </div>
      </div>
      <button class="modal-close is-large" aria-label="close"></button>
    </div>`
})
Vue.component('prod-modal', {
    data: function() {
      return {
        active: false,
        item: {
            product_name: '',
            link: "https://www.foot.com/wp-content/uploads/2017/03/placeholder.gif",
            image: "https://www.foot.com/wp-content/uploads/2017/03/placeholder.gif",
            harga: 0,
            kategori: '',
        }
      }
    },
    methods: {
        show() {
            this.active = true;
        },
        hide() {
            this.active = false;
        }
    },
    created() {
        Event.$on('toggleProdukModal', (row) => {
          console.log(row);
          this.item = row;
          this.show();
        });
    },
    template: `<div class="modal" :class="{ 'is-active': active}">
      <div class="modal-background" v-on:click='hide()'></div>
      <div class="modal-content">
        <div class="card">
                <div class="card-image">
                    <figure class="image is-4by3">
                    <img :src="item.image">
                    </figure>
                </div>
                <div class="card-content">
                    <div class="media">
                        <div class="media-content">
                            <p class="title is-4"><a :href="item.link">{{item.product_name}}</a> - {{item.kategori}}</p>
                            <b><h1>Rp{{item.harga}}</h1></b>
                        </div>
                    </div>
                </div>
        </div>
        </div>
      </div>
      <button class="modal-close is-large" aria-label="close"></button>
    </div>`
});

var app = new Vue({
  el: '#app',
  data() {
    return {
      orders: null,
      produkModalToggle: false,
      pembeliModalToggle: false,
      produkSelected: {
          product_name: '',
          
      },
      pembeliSelected: {
          nama: '',
          email: '',
          alamat: '',
          telepon: '',
      }
    }
  },
  methods: {
    fetchData() {
      this.orders = null;
      axios.get('/orders')
        .then(results => results.data)
        .then(data => (this.orders = data.orders))
        .catch(err => console.log(err))
    },
    selectProduk: function(id) {
        this.produkSelected = this.orders.filter(item => item.id == id);
        Event.$emit('toggleProdukModal', this.produkSelected[0].produk);
    },
    selectPerson: function(id) {
        this.pembeliSelected = this.orders.filter(item => item.id == id);
        Event.$emit('togglePembeliModal', this.pembeliSelected[0].pembeli);
    },
    updateStatus: function(id) {
        Event.$emit('toggleUpdateModal', id);
    }
  },
  filters: {
    moment: (date) => (moment(date).format('Do MMM YYYY, h:mm:ss'))
  },
  created() {
    Event.$on('update', () => {
      console.log('update!');
      this.fetchData();
    });
  },
  mounted() {
    // this.fetchData();
    axios.get('/orders')
      .then(results => results.data)
      .then(data => (this.orders = data.orders))
      .catch(err => console.log(err))
  }
})