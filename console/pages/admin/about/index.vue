<template>
  <div>
    <div class="card">
      <div class="card__body">
        <div class="fn__flex admin__about">
          <div class="about__side">
            <img src="~assets/images/logo.png"/>
          </div>
          <div class="fn__flex-1 vditor-reset">
            <h2 class="fn__clear" v-if="isLatest">
            <span class="fn__left">
              {{ $t('about1', $store.state.locale) }}
              <span>{{ version }}</span>
            </span>
            </h2>
            <h2 class="fn__clear" v-else>
            <span class="fn__left">
               {{ $t('about2', $store.state.locale) }}
              <span class="ft__danger">{{ version }}</span>
            </span>
            </h2>
            <p v-html="$t('about4', $store.state.locale)"></p>
            <p v-html="$t('about3', $store.state.locale)"></p>
            <br>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    data () {
      return {
        isLatest: true,
        download: '',
        version: '',
        list: [],
      }
    },
    head () {
      return {
        title: `${this.$t('about', this.$store.state.locale)} - ${this.$store.state.blogTitle}`,
      }
    },
    async mounted () {
      this.$set(this, 'isLatest', true)
      this.$set(this, 'version', this.$store.state.version)
    },
  }
</script>

<style lang="sass">
  @import '~assets/scss/_variables'

  .admin__about
    .vditor-reset
      padding-left: 50px

      h2
        font-size: 1.4em

    .about
      &__side
        margin-left: 30px
        width: 162px
        align-self: center

      &__github
        border: 0
        margin-left: 20px

      &__link
        margin: 10px 40px 10px 0

        .icon
          border-right: 1px dotted #fff
          padding-right: 10px
          margin: 0 10px 0 0
          height: 18px
          width: 18px

  @media (max-width: 768px)
    .admin__about
      display: block

      .about__github
        float: none
        margin: 10px 0 0 0

      .about__side
        text-align: center
        margin: 15px
        width: auto

      .vditor-reset h2
        font-size: 1.2em
</style>
