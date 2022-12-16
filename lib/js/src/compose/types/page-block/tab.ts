import { PageBlock, PageBlockInput, Registry } from './base'
import { Apply } from '../../../cast'

const kind = 'Tab'

interface Style {
  appearance: string;
  alignment: string;
  fillJustify: string;
  verticalHorizontal: string;
  tabPosition: string;
}
interface Options {
  tabStyle: Style;
  tabs: any[];
  blockIndex: any;
}

const defaults: Readonly<Options> = Object.freeze({
  tabStyle: {
    appearance: 'tabs',
    alignment: 'start',
    fillJustify: '',
    verticalHorizontal: '',
    tabPosition: '',
  },
  tabs: [],
  blockIndex: null,
})

export class PageBlockTab extends PageBlock {
  readonly kind = kind

  options: Options = { ...defaults }

  constructor(i?: PageBlockInput) {
    super(i)
    this.applyOptions(i?.options as Partial<Options>)
  }

  applyOptions(o?: Partial<Options>): void {
    if (!o) return

    Apply(this.options, o, Array, 'tabs')
    Apply(this.options, o, Number, 'blockIndex')
    if (o.tabs) {
      this.options.tabs = o.tabs
    }

    if (o.tabStyle) {
      this.options.tabStyle = o.tabStyle
    }

  }
}

Registry.set(kind, PageBlockTab)
