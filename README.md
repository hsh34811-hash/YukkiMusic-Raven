<div align="center">

# 🎵 YukkiMusic — ✘ RAVEN Edition (2026 Go Engine)

<img src="https://readme-typing-svg.demolab.com?font=Fira+Code&size=30&duration=3000&pause=1000&color=00FFCC&center=true&vCenter=true&width=650&lines=%E2%9C%98+RAVEN+Edition+2026;High-Performance+Go+Engine;Ultra+Low+Memory+(~30MB+RAM);Autonomous+Self-Update+System;YouTube+Anti-Block+Cookies+Support" alt="Typing SVG" />

<p align="center">
  <img src="https://telegra.ph/file/91533956c91d0fd7c9f20.jpg" width="220" height="220" style="border-radius: 50%;" />
</p>

[![Go Version](https://img.shields.io/badge/Go-1.26+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://go.dev/)
[![Telegram](https://img.shields.io/badge/Telegram-Channel-blue?style=for-the-badge&logo=telegram)](https://t.me/Raven_xx24)
[![Docker](https://img.shields.io/badge/Docker-Ready-2496ED?style=for-the-badge&logo=docker&logoColor=white)](https://www.docker.com/)
[![License](https://img.shields.io/badge/License-GPL%20v3-green?style=for-the-badge)](LICENSE)
[![Status](https://img.shields.io/badge/Status-Active%20%26%20Maintained%202026-brightgreen?style=for-the-badge)](https://github.com/hsh34811-hash/YukkiMusic-Raven)

**An ultra-high-performance, next-generation Telegram Group Voice Chat Music & Video Streaming Bot.**  
*Re-engineered from the ground up in Go (Golang) for sub-second latency, zero memory leaks, and autonomous in-chat updates.*

[English Overview](#-english-documentation) | [التوثيق باللغة العربية](#-التوثيق-باللغة-العربية)

</div>

---

# 🇬🇧 English Documentation

## ⚡ Why 2026 Go Engine? (Raven Edition)

Legacy Python-based music bots suffer from heavy memory consumption (300MB–600MB+), frequent WebRTC audio desync, and persistent YouTube IP bans. **YukkiMusic Raven Edition** solves these problems fundamentally:

| Feature | Legacy Python Bots | YukkiMusic Raven Edition (Go) |
| :--- | :--- | :--- |
| **Memory Footprint (RAM)** | ~350MB – 600MB | **~30MB – 50MB (90% reduction)** ⚡ |
| **Startup Time** | 30 – 45 seconds | **< 1 second (Compiled Binary)** 🚀 |
| **WebRTC Voice Engine** | Py-TgCalls (Python wrapper) | **Native `ntgcalls` C++/Go Core** 🔊 |
| **YouTube Anti-Block** | Fails with HTTP 403 / 429 | **Full `COOKIES_LINK` Remote Auto-Loader** 🍪 |
| **In-Chat Self-Update** | Often breaks dependencies | **Autonomous `/update` (rebuilds & hot-restarts)** 🔄 |
| **Memory Leaks** | Frequent over long uptimes | **Zero memory leaks, rock-solid stability** 🛡️ |

---

## 🚀 Key Features

- **High-Fidelity Audio & Video**: Crystal-clear stereo sound and video streaming in Telegram voice chats.
- **Multi-Platform Source Support**: YouTube, YouTube Music, Spotify (Tracks/Playlists/Albums), SoundCloud, and direct Telegram audio/video files.
- **Autonomous In-Chat `/update`**: Type `/update` in Telegram; the bot automatically pulls the latest commits from GitHub, recompiles its binary, and hot-restarts with zero manual terminal work.
- **Remote YouTube Cookies**: Provide a remote link (`COOKIES_LINK`) via Batbin to bypass YouTube bot detection without exposing credentials in Git.
- **Bilingual Core**: Full native support for English (`en.yml`) and Arabic (`ar.yml`) out of the box.

---

## 🛠️ Autonomous Self-Update Feature (`/update`)

Forget logging into your VPS or cloud console to update your bot. As the bot owner, simply send:

```text
/update
```

**What happens behind the scenes:**
1. Bot queries the upstream repository (`https://github.com/hsh34811-hash/YukkiMusic-Raven.git`).
2. Displays incoming commit count and changelog summary directly in Telegram.
3. Automatically executes `git pull`.
4. Rebuilds the Go binary (`go build`).
5. Seamlessly executes `syscall.Exec` hot-restart, re-notifying active rooms.

---

## 🍪 YouTube Anti-Block (Cookies Setup)

YouTube blocks cloud server IPs by default. To make YouTube streaming work 100% reliably:

1. Install browser extension **[Get cookies.txt LOCALLY](https://chromewebstore.google.com/detail/get-cookiestxt-locally)** in your Chrome/Firefox.
2. Visit [YouTube.com](https://www.youtube.com), log into a dummy/secondary Google account.
3. Click the extension and export cookies in **Netscape** format.
4. Go to **[batbin.me](https://batbin.me)**, paste your cookie text, and click **Save**.
5. Copy the generated paste URL and set it in your environment:
   ```env
   COOKIES_LINK="https://batbin.me/your_paste_id"
   ```

---

## 📦 Deployment Guide

### Option 1: Docker (Recommended)

```bash
# Clone the repository
git clone https://github.com/hsh34811-hash/YukkiMusic-Raven.git
cd YukkiMusic-Raven

# Configure environment variables
cp sample.env .env
nano .env

# Build and run with Docker Compose
docker compose up -d --build
```

### Option 2: Linux VPS / Local Host

```bash
# Clone the repository
git clone https://github.com/hsh34811-hash/YukkiMusic-Raven.git
cd YukkiMusic-Raven

# Configure environment variables
cp sample.env .env
nano .env

# Run automated installation & build script
chmod +x install.sh
./install.sh

# Run the compiled binary
./app
```

---

## ⚙️ Environment Variables Reference

| Variable | Required | Description |
| :--- | :---: | :--- |
| `API_ID` | **Yes** | Telegram API ID from [my.telegram.org](https://my.telegram.org) |
| `API_HASH` | **Yes** | Telegram API Hash from [my.telegram.org](https://my.telegram.org) |
| `TOKEN` | **Yes** | Telegram Bot Token from [@BotFather](https://t.me/BotFather) |
| `MONGO_DB_URI` | **Yes** | MongoDB Atlas Connection String |
| `STRING_SESSIONS` | **Yes** | Pyrogram String Session for assistant account |
| `OWNER_ID` | **Yes** | Your Telegram User ID (e.g. from `@userinfobot`) |
| `UPSTREAM_REPO` | No | Git repository URL for `/update` (Defaults to Raven Edition) |
| `UPSTREAM_BRANCH` | No | Git branch for updates (Default: `main`) |
| `COOKIES_LINK` | No | Remote Batbin URL for YouTube authentication |
| `DEFAULT_LANG` | No | Default language (`ar` or `en`, default: `ar`) |
| `SUPPORT_CHAT` | No | Telegram Support link (Default: `https://t.me/Raven_xx24`) |

---

<br>

# 🇸🇦 التوثيق باللغة العربية

## 🌟 لماذا محرك Go لعام 2026؟ (نسخة Raven المطورة)

بوتات الموسيقى التقليدية المكتوبة ببايثون أصبحت تعاني في خوادم تيليجرام الحديثة من استهلاك جنوني للرام (يصل إلى 600 ميجابايت)، وانهيارات مفاجئة في الصوت، وحظر مستمر من يوتيوب.  
تمت إعادة كتابة **YukkiMusic - Raven Edition** بلغة **Go (Golang)** لتقديم أداء استثنائي:

- ⚡ **خفيف كالريشة:** يستهلك فقط **30 إلى 50 ميجابايت رام** (توفير 90% من استهلاك السيرفر).
- 🚀 **إقلاع فوري:** يعمل البوت في أقل من ثانية واحدة بفضل تجميعه كملف ثنائي مستقل.
- 🔄 **تحديث تلقائي ذكي (`/update`):** يحدث البوت نفسه ويستبدل الكود القديم بالجديد ويعيد تشغيل نفسه تلقائياً دون لمس السيرفر.
- 🍪 **تخطي حظر يوتيوب:** دعم مدمج للكوكيز عن بُعد عبر رابط Batbin لتشغيل الموسيقى بدون انقطاع.
- 🌐 **تعريب كامل وشامل:** واجهة البوت والأزرار والرسائل مدعومة باللغة العربية بنسبة 100%.

---

## 🎮 قائمة الأوامر الرئيسية

### 🎵 أوامر التشغيل العامة (للأعضاء والمشرفين):
- `/play [اسم الأغنية أو الرابط]` — تشغيل مقطع صوتي في المكالمة.
- `/vplay [اسم الفيديو أو الرابط]` — تشغيل بث مرئي (فيديو) في المكالمة.
- `/pause` — إيقاف التشغيل مؤقتاً.
- `/resume` — استئناف التشغيل.
- `/skip` — تخطي الأغنية الحالية والانتقال للتالية.
- `/stop` أو `/end` — إيقاف التشغيل بالكامل ومغادرة المكالمة.
- `/seek [الزمن بالثواني]` — تقديم أو تأخير المقطع (مثال: `/seek 30`).
- `/speed [0.5 - 2.0]` — تسريع أو تبطيء سرعة الصوت.
- `/loop [1-10]` — تكرار تشغيل الأغنية الحالية لعدد محدد.
- `/queue` — عرض قائمة الانتظار الحالية.
- `/lang` — تغيير لغة البوت داخل المجموعة.

### 👑 أوامر المالك والمطور (Owner Commands):
- `/update` — **التحديث الذكي:** يفحص كود المستودع الجديد ويسحبه ويعيد بناء البوت وتشغيله ذاتياً.
- `/restart` — إعادة تشغيل عملية البوت ومسح الذاكرة المؤقتة.
- `/cleanmode` — تفعيل وضع الحذف التلقائي للرسائل لمنع تراكم المحادثات.
- `/auth [بالرد]` — ترقية عضو كمشرف صوتيات مخصص.
- `/broadcast [بالرد]` — إذاعة رسالة لجميع المجموعات المشتركة.
- `/speedtest` — قياس سرعة اتصال وإنترنت السيرفر.

---

## 🍪 إعداد كوكيز يوتيوب لتفادي الحظر

1. ثبت إضافة **Get cookies.txt LOCALLY** على متصفحك (Chrome أو Firefox).
2. افتح موقع **YouTube** وسجل بحساب جيميل عادي.
3. اضغط على الإضافة ونزل الكوكيز بصيغة Netscape.
4. افتح موقع **[batbin.me](https://batbin.me)** والصق محتوى الكوكيز واضغط حفظ.
5. انسخ الرابط وضعه في ملف `.env` أمام:
   ```env
   COOKIES_LINK="https://batbin.me/paste_id"
   ```

---

## 👤 المطور والحقوق (Credits)

* **Modified & Enhanced by:** [✘ RAVEN](https://t.me/Raven_xx24)
* **Telegram Channel:** [@Raven_xx24](https://t.me/Raven_xx24)
* **Based on:** TheTeamVivek YukkiMusic Go Core
* **License:** [GNU General Public License v3.0](LICENSE)
