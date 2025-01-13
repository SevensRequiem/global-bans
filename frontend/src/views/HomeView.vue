<script setup lang="ts">
</script>

<template>
  <main>

    <section class="hero">
      <div class="hero-inner">
        <svg width="100px" height="100px" viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" aria-hidden="true" role="img" class="iconify iconify--gis" preserveAspectRatio="xMidYMid meet"><path d="M49.947 0C22.354.03 0 22.406 0 50c0 27.614 22.386 50 50 50s50-22.386 50-50a50 50 0 0 0-2.791-16.307a42.95 42.95 0 0 1-3.508 5.577a44.944 44.944 0 0 1 1.227 8.23h-9.942a28.887 28.887 0 0 1-7.552 3.418a26.037 26.037 0 0 1-8.385-3.418H52.5V30.87h3.014a55.57 55.57 0 0 1-1.965-5H52.5v-3.448C50.594 16.899 49.985 9.017 50 0zM47.5 5.879v19.99H30.75c.988-2.823 2.13-5.429 3.408-7.75C37.89 11.341 42.571 7.102 47.5 5.88zM35.98 7.232c-2.324 2.352-4.41 5.22-6.203 8.475c-1.68 3.05-3.125 6.467-4.312 10.162H12.01c5.535-8.706 13.975-15.37 23.97-18.637zM9.256 30.87h14.808c-1.245 5.162-2.008 10.76-2.203 16.631H5.072a44.79 44.79 0 0 1 4.184-16.63zm19.974 0H47.5V47.5H26.867c.212-5.935 1.043-11.554 2.363-16.63zM5.072 52.5h16.762c.129 5.856.82 11.454 1.994 16.63H9.256A44.79 44.79 0 0 1 5.072 52.5zm21.762 0H47.5v16.63H28.98c-1.245-5.1-2.006-10.715-2.146-16.63zm25.666 0h21.592c-.14 5.915-.902 11.53-2.147 16.63H52.5zm26.576 0h15.852a44.79 44.79 0 0 1-4.184 16.63H77.09c1.17-5.177 1.857-10.775 1.986-16.63zM12.01 74.13h13.136c1.242 4.085 2.8 7.84 4.631 11.165c1.438 2.61 3.068 4.969 4.854 7.017c-9.407-3.41-17.336-9.869-22.621-18.181zm18.394 0H47.5v20.798c-.308-.017-.612-.048-.918-.07c-4.59-1.5-8.924-5.62-12.424-11.975c-1.428-2.594-2.692-5.537-3.754-8.752zm22.096 0h18.021c-1.06 3.216-2.325 6.159-3.753 8.753c-3.428 6.225-7.656 10.308-12.141 11.883a45.48 45.48 0 0 1-2.127.162zm23.275 0H87.99a45.06 45.06 0 0 1-21.228 17.641c1.604-1.92 3.075-4.094 4.386-6.476c1.831-3.325 3.388-7.08 4.627-11.164z" fill="#FFFFFF"></path><path d="M56.5 0A2.5 2.5 0 0 0 54 2.523c.18 18.47 5.5 29.632 10.975 36.096c5.475 6.464 11.312 8.277 11.312 8.277a2.5 2.5 0 0 0 .713.102a2.5 2.5 0 0 0 .027.002a2.5 2.5 0 0 0 .703-.11s5.905-1.87 11.424-8.367C94.674 32.027 100 20.85 100 2.5A2.5 2.5 0 0 0 97.5 0h-41zm2.703 5H77v36.76l-.006.002c-.918-.321-4.14-1.577-8.205-6.375C64.353 30.15 59.834 20.826 59.203 5z" fill="#D64045"></path></svg>
        <h1>Global Bans</h1>
        <p>Protect your game server from malicious users.</p>
        <span>Minecraft</span> | <span>Source</span> | <span>Fail2ban</span>
        <div class="hero-stats">
          <span><font-awesome-icon :icon="['fas', 'server']" /> 0</span>
        <span><font-awesome-icon :icon="['fas', 'ban']" /> 0</span>
        </div>
      </div>
    </section>
    
    <section id="features">
      <div class="wrapper">
        <div class="feature-wrapper">
          <div class="feature">
            <h3><font-awesome-icon :icon="['fas', 'shield-alt']" /> Protection</h3>
            <p>Protect your server from malicious users.</p>
          </div>
          <div class="feature">
            <h3><font-awesome-icon :icon="['fas', 'user-shield']" /> Moderation</h3>
            <p>Keep your server safe with moderation tools.</p>
          </div>
          <div class="feature">
            <h3><font-awesome-icon :icon="['fas', 'user-lock']" /> Bans</h3>
            <p>View and manage bans on your servers.</p>
          </div>
        </div>
      </div>
    </section>

    <section id="config">
      <div class="left">
        <h2>Configuration</h2>
        <p>Configure your server to use Global Bans.</p>
        <p>Download the plugin for your game server.</p>
        <a href="" class="button">Download</a>
      </div>
      <div class="right">
          <span>The URL of the Global Bans Instance.</span>
          <code>url: https://yourdomain.com</code>
          <span>The ID of your server, it is automatically generated when you launch a server with Global Bans and can not be changed.</span>
          <code>serverID: 00000000-0000-0000-0000-000000000000</code>

      </div>
    </section>

    <section id="about">
      <div class="about">
        <h1><font-awesome-icon :icon="['fas', 'info']" /> About</h1>
        <p>Global Bans is a service that allows you to protect your game server from malicious users. It is a free and open-source project that is easy to use and can be set up in minutes.</p>
        <p>Global Bans is compatible with many game servers, including Minecraft, Source, and even the service Fail2ban.</p>
      </div>
    </section>




  </main>
</template>

<style scoped>

.hero {
  background: linear-gradient(180deg, var(--tertiary-color) 10%, var(--background-color) 90%);
  height: 50vh;
  display: flex;
  justify-content: center;
  align-items: center;
  color: var(--text-color);
}

.hero h1 {
  font-size: 3em;
}

.hero p {
  font-size: 1.5em;
  text-align: center;
}

.hero span {
  font-size: 1.5em;
  margin: 0 10px;
}

.hero-stats {
  font-size: 1.5em;
  margin-top: 20px;
}

.hero-stats span {
  margin: 0 10px;
}

#features h2 {
  text-align: center;
}

.wrapper {
  max-width: 1200px;
  margin: 0 auto;
}

.feature-wrapper {
  display: flex;
  justify-content: space-between;
  gap: 20px;
}

.feature {
  flex: 1;
  padding: 20px;
  border: 1px double;
  border-image: linear-gradient(180deg, var(--tertiary-color), var(--background-color)) 1;
  font-size: 1.5em;
  text-align: center;
}

.feature h3 {
  font-size: 1.5em;
  margin-bottom: 10px;
  color: var(--primary-color);
}

#config {
  display: flex;
  justify-content: space-between;
  gap: 20px;
  margin: 40px;
}

#config h2 {
  font-size: 2em;
}

.left {
  flex: 1;
}

.right {
  display: flex;
  flex-direction: column;
  gap: 20px;
  flex: 1;
}

code {
  background: var(--background-color);
  padding: 10px;
  border: 1px solid var(--accent-color);
  border-radius: 5px;

}

.about {
  max-width: 800px;
  margin: 0 auto;
  text-align: center;
}

.about h1 {
  font-size: 2em;
}

.about p {
  font-size: 1.5em;
}

@media (max-width: 768px) {
  .feature-wrapper {
    flex-direction: column;
  }
}

@media (min-width: 1024px) {
  .feature-wrapper {
    gap: 40px;
  }
}



</style>
<script lang="ts">
import { onMounted } from 'vue'
import axios from 'axios'

onMounted(async () => {
  try {
    const response = await axios.get('/api/stats')
    const data = response.data
    if (data.servers && data.bans) {
      // @ts-ignore
      document.querySelector('.hero-stats span:first-child').innerHTML = `<font-awesome-icon :icon="['fas', 'server']" /> ${data.servers} Servers`
      // @ts-ignore
      document.querySelector('.hero-stats span:last-child').innerHTML = `<font-awesome-icon :icon="['fas', 'ban']" /> ${data.bans} Bans`

    }
  } catch (error) {
    console.error('Error fetching stats data:', error)
  }
})

</script>