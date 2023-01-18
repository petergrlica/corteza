import { PageBlock, PageBlockInput, Registry } from './base'
import { Apply } from '../../../cast'

const kind = 'Tabs'

interface Style {
  appearance: string;
  alignment: string;
  fillJustify: string;
  verticalHorizontal: string;
}

interface Tabs {
  block: object;
  indexOnMain: string | number;
}

interface Options {
  style: Style;
  tabs: Tabs[];
  blockIndex: any;
}

const defaults: Readonly<Options> = Object.freeze({
  style: {
    appearance: 'tabs',
    alignment: 'left',
    fillJustify: 'none',
    verticalHorizontal: 'none',
  },
  tabs: [],
  blockIndex: null,
})

export class PageBlockTab extends PageBlock {
  readonly kind = kind

  options: Options = { ...defaults }

  constructor (i?: PageBlockInput) {
    super(i)
    this.applyOptions(i?.options as Partial<Options>)
  }

  applyOptions (o?: Partial<Options>): void {
    if (!o) return

    Apply(this.options, o, Array, 'tabs')
    Apply(this.options, o, Number, 'blockIndex')
    if (o.tabs) {
      this.options.tabs = o.tabs
    }

    if (o.style) {
      this.options.style = o.style
    }
  }
}

Registry.set(kind, PageBlockTab)
