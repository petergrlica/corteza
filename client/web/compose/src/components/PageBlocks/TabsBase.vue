<template>
  <wrap
    :id="`tab-${options.blockIndex}`"
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
        @click="updateBlock(index)"
      >
        <page-block-tab
          v-if="tab.block.kind !== 'Tab'"
          :key="key"
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
      key: 0,
      activeTab: 0,
    }
  },

  watch: {
    '$route.query': {
      handler (query) {
        this.changeActiveTab(query)
      },
    },
  },

  created () {
    const { tabIndex, tabBlockIndex } = this.$route.query
    this.changeActiveTab({ tabBlockIndex, tabIndex })
  },

  methods: {
    // Because some of the blocks are not reactive, we need to force them to update
    updateBlock (index) {
      this.key++
      if (this.$route.query.tabIndex === index && this.$route.query.tabBlockIndex === this.options.blockIndex) {
        return
      }
      this.$router.push({ path: this.$route.fullPath, query: { tabBlockIndex: this.options.blockIndex, tabIndex: index } })
    },

    changeActiveTab (query) {
      if (+query.tabBlockIndex === this.options.blockIndex) {
        this.activeTab = +query.tabIndex
      }
    },
  },
}
</script>
