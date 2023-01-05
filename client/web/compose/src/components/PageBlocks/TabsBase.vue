<template>
  <wrap
    v-bind="$props"
    v-on="$listeners"
  >
    <b-tabs
      v-if="options.tabs.length"
      v-model="activeTab"
      v-bind="{
        align: options.style.alignment,
        fill: options.style.fillJustify === 'fill',
        justified: options.style.fillJustify === 'justified',
        pills: options.style.appearance === 'pills',
        tabs: options.style.appearance === 'tabs',
        small: options.style.appearance === 'small',
        vertical: options.style.verticalHorizontal === 'vertical',
      }"
    >
      <b-tab
        v-for="(tab, index) in options.tabs"
        :key="index"
        :title="tab.block.title"
        @click="trackTab(index)"
      >
        <page-block-tab
          v-if="index === activeTab"
          v-bind="{ ...$attrs, ...$props, page, block: compose().PageBlockMaker(tab.block), blockIndex: index }"
          :record="record"
          :module="module"
        />
      </b-tab>
    </b-tabs>

    <div
      v-else
      class="d-flex justify-content-center pt-5"
    >
      <p class="text-secondary">
        {{ $t('tab.noTabsBase') }}
      </p>
    </div>
  </wrap>
</template>

<script>
import base from './base'
import { compose } from '@cortezaproject/corteza-js'

export default {
  i18nOptions: {
    namespaces: 'block',
  },

  name: 'TabBase',

  components: {
    PageBlockTab: () => import('corteza-webapp-compose/src/components/PageBlocks'),
  },
  extends: base,

  data () {
    return {
      compose: () => compose,
      activeTab: 0,
    }
  },

  watch: {
    '$route.query': {
      immediate: true,

      handler (query) {
        this.changeActiveTab(query)
      },
    },
  },

  methods: {
    trackTab (index) {
      this.$router.replace({ query: { tabBlockIndex: this.options.blockIndex, tabIndex: index } })
    },

    changeActiveTab (query) {
      if (+query.tabBlockIndex === this.options.blockIndex) {
        this.activeTab = +query.tabIndex
      }
    },
  },
}
</script>
