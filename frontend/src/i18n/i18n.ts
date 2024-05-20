import { createI18n } from 'vue-i18n';

import elementde from 'element-plus/es/locale/lang/de';
import elementEn from 'element-plus/es/locale/lang/en';
import elementFr from 'element-plus/es/locale/lang/fr';
import elementJa from 'element-plus/es/locale/lang/ja';
import elementRu from 'element-plus/es/locale/lang/ru';
import elementZh from 'element-plus/es/locale/lang/zh-cn';

import de from './lang/de';
import en from './lang/en';
import fr from './lang/fr';
import ja from './lang/ja';
import ru from './lang/ru';
import zh from './lang/zh';

type ElementPlusLocaleConfig = {
  [key: string]: any;
};

export const elementPlusLocales: ElementPlusLocaleConfig = {
  en: elementEn,
  zh: elementZh,
};

type MessageSchema = typeof en; 


const i18n = createI18n<[MessageSchema], 'de' | 'en' | 'fr' | 'ja' | 'ru' | 'zh'>({
  locale: 'en', 
  messages: {
    'de': { ...elementde, ...de }, 
    'en': { ...elementEn, ...en }, 
    'fr': { ...elementFr, ...fr }, 
    'ja': { ...elementJa, ...ja }, 
    'ru': { ...elementRu, ...ru }, 
    'zh': { ...elementZh, ...zh }, 
  },
});

export default i18n;
