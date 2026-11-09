# Design System — Amerta Technology

## 1. Design Direction

### Design Philosophy

Menggunakan pendekatan **Modern Enterprise Dashboard** dengan karakteristik:

- Professional: memberikan kesan perusahaan teknologi yang kredibel.
- Clean: mengutamakan background putih dan tata letak yang tidak ramai.
- Minimalist: menghindari dekorasi yang tidak memiliki fungsi.
- Consistent: seluruh halaman menggunakan sistem desain dan komponen yang seragam.
- Functional: mengutamakan keterbacaan, efisiensi, dan kemudahan navigasi.

Referensi utama desain adalah identitas visual logo Amerta Technology.

### Visual Direction

- Dominant white background.
- Navy blue sebagai warna utama.
- Teal dan gradient sebagai aksen identitas.
- Light gray untuk pemisah area dan background sekunder.
- Subtle border dan shadow.
- Rounded corners dengan radius moderat.
- Spacious layout dengan whitespace yang cukup.

Hindari desain yang terlalu dekoratif, penggunaan gradient berlebihan, glassmorphism, dan warna-warna mencolok.

---

## 2. Brand Color System

Palet warna terinspirasi dari logo Amerta Technology.

| Token | Color | Usage |
|---|---|---|
| Primary | `#25447F` | Primary buttons, active navigation, links |
| Primary Dark | `#19345F` | Hover states, emphasis |
| Teal | `#42B49B` | Accent, highlights |
| Gradient Start | `#42B49B` | Gradient accent |
| Gradient End | `#24558A` | Gradient accent |
| Background | `#FFFFFF` | Main application background |
| Surface | `#FFFFFF` | Cards and panels |
| Surface Secondary | `#F7F9FC` | Secondary sections |
| Border | `#E8ECF2` | Dividers and component borders |
| Text Primary | `#202B3C` | Headings and primary text |
| Text Secondary | `#667085` | Supporting text |
| Text Muted | `#98A2B3` | Placeholder and disabled text |

### Gradient

Gunakan gradient hanya untuk elemen tertentu.

```css
background: linear-gradient(
  120deg,
  #42B49B 0%,
  #24558A 100%
);
```

Penggunaan:
- Brand accents.
- Small decorative elements.
- Selected promotional components.
- Logo-related visual elements.

Jangan menggunakan gradient sebagai background seluruh dashboard.

### Semantic Colors

| Status | Color |
|---|---|
| Success | `#16A36A` |
| Warning | `#E9A23B` |
| Danger | `#DC5252` |
| Info | `#3182CE` |

Gunakan semantic colors secara konsisten untuk status, notifikasi, dan feedback.

---

## 3. Typography

Gunakan font **Inter** sebagai tipografi utama.

Alternatif:
- Geist
- Plus Jakarta Sans

### Typography Scale

| Element | Size | Weight |
|---|---:|---|
| Page Title | 24px | 600 |
| Section Title | 18px | 600 |
| Card Title | 15px | 600 |
| Body | 14px | 400 |
| Secondary Text | 13px | 400 |
| Caption | 12px | 400 |
| Button | 14px | 500 |

### Typography Rules

- Gunakan warna gelap untuk heading.
- Hindari penggunaan terlalu banyak font weight.
- Pastikan kontras teks memadai.
- Jangan menggunakan font size terlalu kecil untuk informasi penting.
- Gunakan line-height yang nyaman dibaca.

---

## 4. Application Layout

Dashboard menggunakan struktur utama:

```text
+--------------------------------------------------+
| Sidebar     | Top Navigation                     |
|             +------------------------------------+
|             |                                    |
| Navigation  | Page Header                        |
|             |                                    |
|             | Main Content                       |
|             |                                    |
|             | Cards / Tables / Charts            |
|             |                                    |
|             |                                    |
+--------------------------------------------------+
```

### Sidebar

Sidebar harus memiliki tampilan sederhana dan profesional.

Specifications:
- Width: 250px.
- Background: #FFFFFF.
- Border-right: 1px solid #E8ECF2.
- Position: Fixed atau sticky.
- Logo ditempatkan di bagian atas.
- Navigation dikelompokkan berdasarkan kategori.
- Active navigation menggunakan background biru sangat muda.
- Icon menggunakan Lucide Icons.
- Collapsible sidebar tersedia jika diperlukan.

Active navigation:

```css
background: #EEF3FA;
color: #25447F;
font-weight: 600;
border-radius: 8px;
```

Hindari sidebar dengan background biru gelap penuh. Pertahankan dominasi warna putih.

### Top Navigation

Specifications:
- Height: 64px.
- Background: #FFFFFF.
- Border-bottom: 1px solid #E8ECF2.
- Search field jika diperlukan.
- Notification icon.
- User profile dan dropdown.
- Breadcrumb untuk halaman bertingkat.

Top navigation harus ringan dan tidak mendominasi konten.

### Main Content

Specifications:
- Background: #FFFFFF atau #F7F9FC.
- Padding desktop: 28px–32px.
- Maximum width mengikuti kebutuhan konten.
- Consistent spacing.
- Responsive layout.

---

## 5. Dashboard Components

### 5.1 Statistic Cards

Digunakan untuk menampilkan informasi ringkas seperti total transaksi, pengguna, pendapatan, dan aktivitas.

Design:
- White background.
- Border: 1px solid #E8ECF2.
- Border radius: 12px.
- Padding: 20px.
- Subtle shadow.
- Icon dengan background berwarna lembut.
- Nilai utama menggunakan font size 24px–28px.

Hindari penggunaan gradient pada setiap card.

### 5.2 Data Tables

Tabel merupakan komponen utama aplikasi dashboard.

Design specifications:
- White background.
- Header background: #F8FAFC.
- Header text: #667085.
- Row height: 52px–60px.
- Border-bottom: 1px solid #E8ECF2.
- Hover background: #F7F9FC.
- Border radius: 10px–12px.

Pada desktop, tabel adalah area scroll utama. Shell mengunci tinggi aplikasi ke
viewport (`lg:h-dvh` + `lg:overflow-hidden`), card daftar berhenti 20px di atas
tepi bawah viewport, dan hanya container tabel yang memiliki scroll internal.
Detail rantai class dan invariant ada di §19.9.

Fitur yang perlu didukung jika relevan:
- Search.
- Sorting.
- Filtering.
- Pagination.
- Row actions.
- Bulk selection.
- Empty state.
- Loading state.

Pastikan tabel tetap nyaman digunakan pada layar kecil melalui horizontal scrolling atau responsive presentation.

### 5.3 Buttons

Primary Button:
- Background: #25447F.
- Text: #FFFFFF.
- Border radius: 8px.
- Height: 40px.
- Font weight: 500.

Secondary Button:
- Background: #FFFFFF.
- Text: #344054.
- Border: 1px solid #D0D5DD.

Ghost Button:
- Transparent background.
- No visible border.
- Text: #475467.

Danger Button:
- Background: #DC5252.
- Text: #FFFFFF.

Button harus memiliki hover, active, focus, disabled, dan loading states.

### 5.4 Form Elements

Input:
- Height: 42px.
- Border: 1px solid #D0D5DD.
- Border radius: 8px.
- Background: #FFFFFF.
- Padding horizontal: 12px.

Focus state:

```css
border-color: #25447F;
box-shadow: 0 0 0 3px rgba(37, 68, 127, 0.12);
```

Ketentuan:
- Label berada di atas input.
- Error message berada di bawah input.
- Required field memiliki indikator.
- Placeholder menggunakan warna muted.
- Gunakan helper text jika diperlukan.
- Select memakai komponen custom `Select` (§20.7), bukan `<select>` native.
  Trigger setinggi 42px mengikuti gaya input; panel listbox dirender portal agar
  tidak terpotong modal, dengan lebar sama dengan trigger dan scrollbar dihitung
  di dalam lebar panel.

### 5.5 Status Badges

Gunakan badge dengan background lembut dan teks yang memiliki kontras memadai.

Contoh:
- Success: light green.
- Pending: light yellow.
- Failed: light red.
- Information: light blue.

Badge harus berukuran proporsional dan tidak terlalu mencolok.

### 5.6 Dialog & Modal

- White background.
- Border radius: 12px.
- Subtle shadow.
- Overlay berwarna hitam dengan opacity rendah.
- Clear title and description.
- Primary dan secondary action dipisahkan secara visual.

Gunakan modal hanya untuk interaksi yang membutuhkan perhatian pengguna.

---

## 6. Spacing & Border Radius

Gunakan sistem spacing berbasis kelipatan 4px.

| Token | Value |
|---|---:|
| XS | 4px |
| SM | 8px |
| MD | 12px |
| LG | 16px |
| XL | 24px |
| 2XL | 32px |
| 3XL | 48px |

Border radius:

| Component | Radius |
|---|---:|
| Input | 8px |
| Button | 8px |
| Card | 12px |
| Modal | 12px |
| Badge | 6px |
| Dropdown | 8px |

Gunakan konsistensi, bukan variasi radius pada setiap komponen.

---

## 7. Iconography

Gunakan Lucide Icons sebagai icon library utama.

Ketentuan:
- Gunakan outline icons.
- Default size: 18px–20px.
- Stroke width: 1.75–2.
- Gunakan icon sesuai konteks.
- Hindari mencampur berbagai icon styles.

Icon tidak boleh menggantikan label untuk navigasi yang membutuhkan kejelasan.

---

## 8. Interaction & Micro-interactions

Interaksi harus terasa responsif tetapi tidak berlebihan.

Standar:
- Transition duration: 150ms–200ms.
- Hover state yang subtle.
- Focus state yang jelas.
- Loading indicator pada proses asynchronous.
- Toast notification untuk feedback.
- Skeleton loading untuk konten yang sedang dimuat.
- Confirmation dialog untuk tindakan destruktif.

Gunakan animasi seperlunya. Hindari animasi dekoratif yang mengganggu produktivitas.

---

## 9. Responsive Design

Dashboard harus mendukung:

- Desktop: 1440px.
- Laptop: 1280px.
- Tablet: 768px.
- Mobile: 375px.

### Desktop
- Sidebar selalu terlihat.
- Main content menggunakan ruang secara optimal.
- Tabel menggunakan lebar yang tersedia.

### Tablet
- Sidebar dapat di-collapse.
- Grid cards menyesuaikan lebar layar.
- Navigasi tetap mudah diakses.

### Mobile
- Sidebar berubah menjadi drawer.
- Top navigation disederhanakan.
- Cards menggunakan satu atau dua kolom sesuai lebar.
- Form menggunakan full width.
- Tabel menggunakan horizontal scroll atau layout alternatif.

Jangan sekadar mengecilkan tampilan desktop. Setiap komponen harus disesuaikan dengan konteks penggunaan mobile.

---

## 10. Accessibility

Ikuti prinsip WCAG 2.2 AA sejauh relevan.

- Pastikan kontras teks memadai.
- Semua input memiliki label.
- Navigasi dapat digunakan dengan keyboard.
- Focus indicator harus terlihat.
- Icon-only buttons memiliki accessible label.
- Jangan hanya mengandalkan warna untuk menunjukkan status.
- Gunakan semantic HTML.
- Pastikan dialog mendukung keyboard interaction.

---

## 11. Design Anti-Patterns

Hindari hal-hal berikut:

1. Dark sidebar dengan gradient yang terlalu dominan.
2. Terlalu banyak warna dalam satu halaman.
3. Card dengan shadow tebal.
4. Border radius terlalu besar pada semua komponen.
5. Glassmorphism berlebihan.
6. Gradient pada setiap tombol dan card.
7. Typography yang terlalu kecil.
8. Spacing yang tidak konsisten.
9. Terlalu banyak elemen dekoratif.
10. Penggunaan emoji sebagai pengganti icon UI.
11. Layout yang tidak konsisten antarhalaman.
12. Menampilkan semua informasi dalam satu halaman tanpa hierarki.

---

## 12. Implementation Guidelines

Design system ini harus diterapkan secara konsisten pada seluruh aplikasi.

Ketentuan implementasi:

- Buat reusable components.
- Gunakan centralized design tokens.
- Hindari hardcoded colors yang tersebar di berbagai komponen.
- Gunakan satu sistem typography.
- Gunakan konsistensi spacing.
- Pisahkan UI components dari business logic.
- Pastikan seluruh komponen memiliki state yang lengkap.
- Pertahankan konsistensi visual ketika menambahkan fitur baru.

### CSS Variables Reference

```css
:root {
  --color-primary: #25447F;
  --color-primary-dark: #19345F;
  --color-teal: #42B49B;

  --color-background: #FFFFFF;
  --color-surface: #FFFFFF;
  --color-surface-secondary: #F7F9FC;

  --color-border: #E8ECF2;

  --color-text-primary: #202B3C;
  --color-text-secondary: #667085;
  --color-text-muted: #98A2B3;

  --color-success: #16A36A;
  --color-warning: #E9A23B;
  --color-danger: #DC5252;
  --color-info: #3182CE;

  --radius-sm: 6px;
  --radius-md: 8px;
  --radius-lg: 12px;

  --spacing-1: 4px;
  --spacing-2: 8px;
  --spacing-3: 12px;
  --spacing-4: 16px;
  --spacing-6: 24px;
  --spacing-8: 32px;
}
```

---

## 13. Final Design Principles

Setiap halaman baru harus mengikuti prinsip berikut:

1. White-first: putih sebagai warna dominan, bukan sekadar background.
2. Navy blue sebagai jangkar identitas, teal sebagai aksen terbatas.
3. Hierarki visual jelas: page title, section title, konten, metadata.
4. Konsisten dengan komponen dan spacing yang sudah ada.
5. Fungsional: setiap elemen memiliki tujuan, bukan dekorasi.
6. Tenang dan profesional: hindari warna, gradient, dan animasi berlebihan.
7. Data-first: angka, status, dan aksi utama mudah dipindai.
8. Aksesibel: kontras, label, focus state, dan navigasi keyboard.
9. Responsif secara kontekstual, bukan hanya mengecilkan desktop.
10. Hemat: tidak menambah library atau pola visual baru tanpa kebutuhan.

Prinsip di atas berlaku untuk dashboard, departemen, karyawan, sesi, audit, serta seluruh state (loading, empty, error).

---

## 14. Brand Identity & Logo Usage

### 14.1 Karakter Brand

Amerta Technology diposisikan sebagai perusahaan teknologi yang:

- profesional,
- terpercaya,
- modern namun tidak trendi berlebihan,
- rapi dan efisien.

Identitas visual diturunkan dari logo: navy blue, teal, dan aksen gradient teal-biru. Dashboard harus terasa seperti produk enterprise yang serius, bukan landing page marketing.

### 14.2 Logo dan Brand Mark

- Gunakan file logo resmi `public/logo.png` sebagai penanda aplikasi di login dan sidebar.
- Logo ditampilkan tanpa teks “Amerta” maupun “Admin Dashboard”; konteks aplikasi tetap tersedia melalui judul halaman dan `title` dokumen.
- Jangan memutar, meregangkan, atau mengganti warna logo.
- Jangan menempatkan logo di atas background ramai atau gradient penuh.
- Di sidebar, logo berada di area atas dengan padding yang konsisten.
- Pertahankan proporsi logo dengan `object-contain` dan sediakan `alt` yang deskriptif.
- Logo bukan pengganti judul halaman; judul halaman tetap ditampilkan di top navigation.

### 14.3 Penggunaan Gradient

Gradient hanya untuk:

- brand mark,
- garis aksen tipis di bagian atas halaman login,
- elemen kecil yang memang menjadi identitas.

Jangan menggunakan gradient untuk:

- background seluruh halaman,
- setiap card,
- setiap tombol,
- sidebar,
- tabel.

### 14.4 Tone Visual

- Sudut membulat moderat (8–12px), bukan pill di semua elemen.
- Shadow tipis, bukan bayangan tebal.
- Border tipis `#E8ECF2` untuk memisahkan area.
- Whitespace cukup; jangan memadatkan semua informasi.
- Ikon outline, bukan ikon berwarna-warni.

---

## 15. Information Architecture & Navigation

### 15.1 Struktur Navigasi

Implementasi saat ini:

```text
Amerta Technology
  ├── Departemen   → /departments
  └── Karyawan     → /employees

Akun (topbar desktop / drawer mobile)
  ├── Hubungi customer service (WhatsApp)
  └── Keluar
```

Target jangka panjang (belum ada):

```text
Dashboard
  └── Overview
Keamanan
  └── Sesi Aktif
Sistem
  └── Audit Log
Akun
  └── Profil (future)
```

### 15.2 Sidebar Items

| Label | Ikon | Route | Status |
|---|---|---|---|
| Departemen | Building | `/departments` | Aktif |
| Karyawan | Users | `/employees` | Aktif |
| Dashboard | `LayoutDashboard` | `/` | Belum ada |
| Sesi Aktif | `MonitorSmartphone` | `/sessions` | Belum ada |
| Audit Log | `ScrollText` | `/audit-logs` | Belum ada |

- Sidebar hanya menampilkan dua menu aktif; tidak ada grup label.
- Active state menggunakan background `#EEF3FA`, teks `#25447F`, weight 600.
- Ikon dibuat inline SVG; tidak memakai library ikon.
- Sidebar tidak menampilkan username; username, customer service, dan logout ada
  di topbar (desktop) atau drawer (mobile).
- Jangan menampilkan menu future sebelum halamannya ada.

### 15.3 Top Navigation

- Breadcrumb: `Dashboard / Departemen / Engineering`.
- Search global tidak wajib pada MVP; jika ditambahkan, jangan menduplikasi search tabel secara membingungkan.
- Notification icon hanya jika ada sumber notifikasi nyata; jangan menampilkan badge palsu.
- Username hanya ditampilkan di sisi kanan top navigation, diikuti icon customer service dan icon logout (desktop).
- Sidebar tidak menampilkan username ataupun aksi customer service/logout pada desktop.
- Mobile: menu button membuka sidebar drawer; breadcrumb disederhanakan.
- Mobile: aksi logout dan customer service ditempatkan di bagian bawah sidebar drawer.

### 15.4 Page Header Pattern

Setiap halaman memiliki:

1. Breadcrumb (kecuali dashboard).
2. Page title (24px, weight 600).
3. Deskripsi singkat (14px, `#667085`) jika membantu.
4. Primary action di kanan (desktop) atau bawah title (mobile).
5. Area filter/aksi sekunder di bawah header.

Jangan menumpuk terlalu banyak tombol. Maksimal satu primary action per halaman.

---

## 16. Dashboard Overview UI

### 16.1 Tujuan

Dashboard harus menjawab pertanyaan berikut dalam < 5 detik:

- Berapa departemen?
- Berapa total karyawan?
- Berapa yang aktif dan nonaktif?
- Bagaimana distribusinya?
- Apa yang baru terjadi?

### 16.2 Layout

```text
Page Header: Dashboard
Subtitle: Ringkasan struktur organisasi Amerta Technology

[Total Departemen] [Total Karyawan] [Karyawan Aktif] [Karyawan Nonaktif]

[Distribusi Karyawan per Departemen]   [Aktivitas Terbaru]
[                                        ]
```

Desktop:

- Stat cards: 4 kolom, gap 16–24px.
- Baris kedua: distribusi (2/3) dan aktivitas (1/3).
- Gunakan grid 12 kolom; jangan membuat card terlalu sempit.

Tablet:

- Stat cards: 2 kolom.
- Distribusi dan aktivitas bertumpuk.

Mobile:

- Stat cards: 1–2 kolom (2 kolom untuk angka pendek).
- Semua card full width.
- Aktivitas menampilkan 5 item pertama; “Lihat semua” menuju audit log jika tersedia.

### 16.3 Statistic Cards

Setiap card:

- Label 13px, `#667085`.
- Nilai 28px, weight 600, `#202B3C`.
- Icon 20px dalam lingkaran lembut.
- Delta/trend hanya jika ada data historis; **jangan** menampilkan panah naik/turun tanpa sumber data.
- Deskripsi singkat opsional (mis. “Seluruh departemen”).
- Card aktif/nonaktif boleh memakai tint success/neutral yang sangat lembut, bukan warna penuh.

Warna icon:

| Card | Tint |
|---|---|
| Total departemen | Navy/blue lembut |
| Total karyawan | Teal lembut |
| Karyawan aktif | Success lembut |
| Karyawan nonaktif | Neutral/abu lembut |

Jangan memakai danger merah untuk karyawan nonaktif; nonaktif adalah status normal, bukan error.

### 16.4 Distribusi Karyawan

MVP tanpa chart library:

- Gunakan horizontal bar list:
  - Nama departemen + kode.
  - Bar dengan track `#F2F4F7`, fill navy/teal.
  - Jumlah dan persentase.
  - Status badge kecil jika departemen nonaktif.
- Urutkan dari jumlah terbesar; departemen 0 tetap tampil di bawah.
- Tampilkan maksimal 8–10 departemen; jika lebih, sediakan “Lihat semua departemen”.
- Jika chart library disetujui, gunakan bar chart sederhana dengan warna navy/teal; jangan pie chart 3D atau warna-warni.
- Selalu sediakan angka tekstual; warna bar bukan satu-satunya penyampai data.

### 16.5 Aktivitas Terbaru

- Maksimal 10 item.
- Setiap item: icon kategori, ringkasan, aktor, waktu relatif.
- Waktu absolut tersedia via tooltip/`title`.
- Aksi mengarah ke detail audit jika viewer tersedia.
- Jika audit belum tersedia/kosong, tampilkan empty state, bukan section yang disembunyikan tanpa penjelasan.
- Jangan menampilkan data sensitif (token, password).

### 16.6 State

- **Loading:** skeleton untuk cards, bar list, dan activity.
- **Error:** card error dengan “Coba lagi”; jangan menampilkan 0.
- **Empty:** jika belum ada data, tampilkan CTA ke departemen/karyawan.
- **Partial:** jika audit gagal tetapi statistik berhasil, tampilkan statistik dan error kecil pada panel aktivitas.

---

## 17. Department Management UI

### 17.1 Department List

Header:

- Title “Departemen”.
- Subtitle “Kelola struktur departemen dan lihat jumlah karyawannya.”
- Primary button “Tambah departemen” (ikon `Plus`).

Toolbar (satu baris di desktop, stacked di mobile):

- Search input dengan icon `Search`, placeholder “Cari kode atau nama
  departemen”. Tombol “Cari” berada di dalam input dan baru menerapkan
  pencarian saat disubmit.
- Status filter select: “Semua status”, “Aktif”, “Nonaktif”. Perubahan langsung
  memuat ulang daftar.
- Tombol “Reset” muncul hanya jika search atau status aktif.
- Tidak ada result count terpisah di toolbar; rentang data ditampilkan di
  pagination desktop atau sentinel mobile.

Desktop data grid (≥ 1024px):

- Grid mengisi sisa tinggi viewport setelah header halaman, toolbar, pagination,
  dan padding bawah; tepi bawah card berhenti **20px di atas tepi bawah
  viewport**.
- Area tabel memakai `flex-1` + scroll internal; halaman tidak ikut memanjang.
- Jika viewport pendek, area tabel menyusut dan tetap dapat di-scroll. Tidak ada
  tinggi tetap 520px.
- Implementasi: Shell memakai `lg:h-dvh lg:overflow-hidden`; `main`, container,
  root halaman, section, dan area tabel memakai `min-h-0`; toolbar dan pagination
  `shrink-0`.
- `table-fixed` + `<colgroup>`; lebar kolom tidak berubah saat filter berubah.
- Header sticky di dalam container scroll.

Table columns:

| Kolom | Lebar | Alignment | Catatan |
|---|---:|---|---|
| Departemen | 26% | Kiri | Nama truncate + kode uppercase. |
| Deskripsi | 34% | Kiri | Satu baris truncate; “—” jika kosong. |
| Karyawan | 14% | Kiri | `{n} karyawan`, tabular-nums. |
| Status | 14% | Kiri | Badge Aktif/Nonaktif. |
| Aksi | 12% | Kanan | Menu tiga titik (`MoreHorizontal`). |

Mobile (< 1024px):

- Bukan tabel; gunakan card list.
- Container section/toolbar transparan (tanpa background, border, shadow); hanya
  card item dan input search yang memakai background putih.
- Setiap card: nama + kode, badge status + menu tiga titik, deskripsi 2 baris,
  footer jumlah karyawan.
- Pagination diganti **infinite scroll** 10 item per halaman dengan
  `IntersectionObserver`.
- Sentinel menampilkan status “Gulir untuk memuat data lainnya”, “Memuat
  data...”, tombol retry, atau “Menampilkan n dari total data”.

Row actions:

- “Lihat karyawan” → `/employees?department={public_id}`.
- “Edit” → modal form.
- “Hapus” (danger; disabled dengan tooltip jika masih memiliki karyawan).

Perilaku:

- Menu aksi memakai portal, menutup saat scroll/resize/Escape/klik overlay.
- Search tidak debounce; memakai submit untuk mengurangi request.
- Sorting belum tersedia di UI; backend mengurutkan nama.
- Pagination desktop: “Sebelumnya”, “{page} / {pages}”, “Berikutnya”, plus
  rentang “{from}-{to} dari {total} data”.
- Setiap perubahan filter mereset daftar ke halaman 1 dan scroll ke atas.

### 17.2 Department Detail

> **Status:** belum diimplementasikan. Saat ini daftar departemen dan filter
> karyawan per departemen sudah tersedia, tetapi halaman detail terpisah belum
> ada.

Layout:

```text
Breadcrumb
[Department name] [status badge]              [Ubah] [Nonaktifkan] [Hapus]
Kode • Deskripsi

[Total karyawan] [Aktif] [Nonaktif] [Dibuat]

[Karyawan di departemen ini]
[Search] [Status filter]                       [Tambah karyawan]
[Table]
```

- Deskripsi kosong ditampilkan sebagai “Belum ada deskripsi.” dengan warna muted.
- Metadata tanggal memakai format `dd MMM yyyy, HH:mm` (locale Indonesia).
- Tombol hapus disembunyikan/disabled jika `employee_count > 0`; tampilkan alasan.
- Tab tidak diperlukan untuk MVP; gunakan section.

### 17.3 Department Form

Form create dan edit memakai modal yang sama:

1. Field **Kode** (wajib, uppercase otomatis, max 20).
2. Field **Nama departemen** (wajib, max 100).
3. Field **Deskripsi** (textarea, max 500).
4. Field **Status** (select Aktif/Nonaktif).
5. Footer: “Batal”, “Simpan”.

- Layout desktop: Kode 140px + Nama fleksibel; deskripsi dan status full width.
- Mobile: satu kolom.
- Error validasi/server tampil di `ErrorAlert` di dalam modal.
- Tombol simpan disabled saat submit dan berlabel “Menyimpan...”.
- Setelah sukses: modal form ditutup, **modal respons** sukses muncul, lalu
  daftar di-refresh.
- Tidak ada toast; semua feedback create/update/delete memakai
  `ResponseModal`.
- Modal dapat ditutup lewat tombol close, overlay, atau Escape.

### 17.4 Delete/Deactivate Dialog

Delete (sudah diimplementasikan):

- Judul: “Hapus departemen?”.
- Body: “Departemen **Engineering** akan dihapus permanen. Tindakan ini tidak
  dapat dibatalkan.”
- Jika `employees > 0`: aksi “Hapus” di menu baris disabled dengan tooltip
  “Pindahkan atau hapus karyawan terlebih dahulu”.
- Tombol: “Batal” (secondary), “Hapus” (danger).
- Wajib konfirmasi; tidak boleh delete langsung dari row.
- Saat proses, tombol disabled dan berlabel “Memproses...”.
- Sukses/gagal ditampilkan di `ResponseModal`; backend juga menolak dengan 409
  jika masih ada karyawan.

Deactivate (belum ada aksi cepat terpisah):

- Status diubah melalui form edit (select Status).
- Target dialog konfirmasi nonaktif tetap berlaku jika aksi cepat ditambahkan
  nanti.

---

## 18. Employee Management UI

### 18.1 Employee List

Header:

- Title “Karyawan”.
- Subtitle “Kelola karyawan dan penempatan mereka di setiap departemen.”
- Primary “Tambah karyawan”; disabled selama opsi departemen belum tersedia.

Toolbar:

- Search “Cari nama, ID, atau email”; diterapkan saat submit.
- Department filter select; nilai opsi adalah `public_id`, dan perubahan
  langsung menulis `/employees?department={public_id}`.
- Status filter: “Semua status”, “Aktif”, “Nonaktif”.
- Reset muncul jika ada search/status/departemen aktif.
- Tidak ada result count terpisah di toolbar.

Desktop data grid (≥ 1024px):

- Grid mengisi sisa tinggi viewport; tepi bawah card berhenti 20px di atas tepi
  bawah viewport.
- Area tabel memakai `flex-1` + scroll internal; halaman tidak memanjang.
- Shell memakai `lg:h-dvh lg:overflow-hidden`; `main`, container, root halaman,
  section, dan area tabel memakai `min-h-0`; toolbar dan pagination `shrink-0`.
- `table-fixed` + `<colgroup>`; total lebar 100%.
- Header sticky.
- Pagination 10 item per halaman.

Table columns:

| Kolom | Lebar | Catatan |
|---|---:|---|
| Karyawan | 22% | Avatar inisial 32px + nama + employee ID. |
| Kontak | 22% | Email truncate + telepon. |
| Departemen | 14% | Nama departemen. |
| Jabatan | 14% | Truncate; “—” jika kosong. |
| Bergabung | 10% | `dd MMM yyyy` locale id-ID. |
| Status | 10% | Badge. |
| Aksi | 8% | Menu tiga titik. |

Mobile (< 1024px):

- Card list dengan avatar inisial, nama, employee ID, badge status, dan menu
  tiga titik.
- Container section/toolbar transparan; hanya card item dan input search yang
  memakai background putih.
- Grid 2 kolom: email (full width), telepon, departemen, jabatan, bergabung.
- Infinite scroll 10 item per halaman + skeleton card.

Row actions:

- “Edit”.
- “Hapus” (danger; disabled dengan tooltip jika status masih aktif).

Perilaku:

- Avatar inisial: dua huruf, background primary lembut, teks primary.
- Email tampil di kolom Kontak, bukan di bawah nama pada grid desktop.
- Sorting belum tersedia di UI; backend mengurutkan `full_name`.
- Jangan menyembunyikan status dan departemen di mobile.

### 18.2 Employee Detail

> **Status:** belum diimplementasikan. Semua data karyawan saat ini dikelola
> lewat daftar + modal edit.

Layout:

```text
Breadcrumb
[Avatar] [Nama] [status badge]                  [Ubah] [Pindahkan] [Nonaktifkan]
[Employee ID] [Email] [Telepon]

[Informasi Kepegawaian]        [Departemen]
Jabatan                        Nama + kode
Departemen                     Jumlah karyawan
Tanggal bergabung              [Lihat departemen]
Status
Dibuat / diperbarui

[Aktivitas terkait] (jika audit tersedia)
```

- Data kosong (phone) tampil “—” dengan warna muted, bukan field hilang.
- Tombol hapus hanya muncul untuk inactive.
- “Pindahkan” membuka dialog.
- Jangan menampilkan data sensitif.

### 18.3 Employee Form

Modal form berisi:

1. **Identitas**
   - Employee ID (wajib, uppercase, max 20; disabled saat edit).
   - Nama lengkap (wajib, max 100).
   - Email (wajib, type email, max 254).
   - Nomor telepon (opsional, max 20).
2. **Kepegawaian**
   - Departemen (select wajib; opsi UI memakai `public_id`, request mengirim
     `department_id` numerik).
   - Jabatan (opsional, max 100).
   - Status (Aktif/Nonaktif).
   - Tanggal bergabung (date).

- Desktop: 2 kolom untuk pasangan field; mobile: 1 kolom.
- Field wajib ditandai `*` dan `aria-required`.
- Saat create dari halaman yang difilter departemen, departemen tersebut
  otomatis terpilih.
- Opsi departemen dimuat dari `GET /api/departments?limit=200`.
- Error validasi/server tampil di dalam modal.
- Sukses: modal ditutup, `ResponseModal` sukses muncul, daftar di-refresh.
- `employee_id` tidak diubah oleh backend saat update.

### 18.4 Transfer Dialog

> **Status:** belum diimplementasikan sebagai dialog terpisah. Perpindahan
> departemen dilakukan melalui select Departemen di form edit.

- Judul “Pindahkan karyawan”.
- Body: nama karyawan, departemen asal → departemen tujuan.
- Select departemen aktif, exclude current.
- Peringatan: “Riwayat perpindahan akan dicatat di audit log.”
- Tombol “Batal”, “Pindahkan”.
- Jika tidak ada departemen aktif lain: empty state di dialog + link untuk membuat/mengaktifkan departemen.
- Setelah sukses: toast, refresh detail dan list.

### 18.5 Status Dialog

> **Status:** belum ada dialog status terpisah. Status diubah lewat select
> Status di form edit.

- Deactivate: “Nonaktifkan karyawan?” → “Karyawan tidak akan dihitung sebagai karyawan aktif. Data tetap tersimpan.”
- Reactivate: “Aktifkan kembali karyawan?”.
- Tombol sesuai aksi; loading state.
- Wajib konfirmasi.

### 18.6 Delete Dialog

- Hanya untuk inactive.
- Judul “Hapus karyawan?”.
- Body menyebut nama dan Employee ID, serta “Tindakan ini permanen.”
- Tombol Batal / Hapus (danger).
- Saat proses, tombol disabled dan berlabel “Memproses...”.
- Jika aktif, aksi hapus di menu baris disabled dengan tooltip “Nonaktifkan
  karyawan terlebih dahulu”.
- Sukses/gagal ditampilkan di `ResponseModal`.

---

## 18A. Session Management UI

### 18A.1 Session List

Header:

- Title “Sesi Aktif”.
- Subtitle “Kelola perangkat dan sesi yang memiliki akses ke akun Anda.”
- Tidak ada primary action; aksi bersifat per sesi.

Table columns:

| Kolom | Catatan |
|---|---|
| Perangkat | User agent yang disederhanakan (browser + OS), bukan string mentah panjang. |
| IP | Alamat IP; bisa masked untuk tampilan. |
| Dibuat | Waktu sesi dibuat. |
| Terakhir dipakai | Waktu refresh terakhir. |
| Kedaluwarsa | Waktu sesi berakhir. |
| Status | Badge “Sesi ini” (info) atau “Aktif” (success). |
| Aksi | “Cabut sesi” (danger/ghost) dengan konfirmasi. |

- Jika pengguna memiliki permission melihat sesi pengguna lain, tampilkan kolom “Pengguna” dan filter pengguna.
- Jangan pernah menampilkan token, hash, cookie, atau family ID mentah di UI.
- Sesi saat ini ditandai jelas.
- Urutkan sesi saat ini di atas, lalu `last_used_at` terbaru.

### 18A.2 Revoke Dialog

- Judul “Cabut sesi?”.
- Body: perangkat, IP, dan waktu terakhir dipakai.
- Jika sesi saat ini: peringatan “Anda akan keluar dari perangkat ini.”
- Tombol “Batal” dan “Cabut sesi” (danger).
- Setelah sukses: toast dan refresh list; jika sesi saat ini, redirect ke login.

### 18A.3 States

- Loading: skeleton rows.
- Empty: “Tidak ada sesi aktif lain.” Jika hanya sesi saat ini, tetap tampilkan sesi tersebut.
- Error: panel retry.
- Revoked: hilang dari daftar aktif atau tampil dengan badge “Dicabut” jika filter menyertakan.

### 18A.4 Mobile

- Card per sesi: perangkat, IP, waktu, status, tombol cabut.
- Dialog full-width.
- Jangan menyembunyikan status “Sesi ini”.

---

## 18B. Audit Log UI

### 18B.1 Audit List

Header:

- Title “Audit Log”.
- Subtitle “Riwayat aktivitas administratif dan keamanan.”
- Read-only; tidak ada tombol create/edit/delete.

Toolbar:

- Filter action (select/search).
- Filter entity type.
- Filter actor.
- Rentang tanggal (`from`, `to`).
- Reset filter.
- Jumlah hasil.

Table columns:

| Kolom | Catatan |
|---|---|
| Waktu | Absolut `dd MMM yyyy, HH:mm:ss`; tooltip relatif. |
| Aktor | Username; “Sistem” untuk aktor null. |
| Aksi | Badge kategori + label manusiawi. |
| Entitas | Tipe + ID. |
| Ringkasan | Satu baris, truncate. |
| IP | Opsional, tampil secondary. |
| Aksi | “Lihat detail”. |

- Urutkan `created_at` terbaru.
- Pagination server-side.
- Gunakan tampilan denser (row 44–52px) tetapi tetap terbaca.
- Kategori badge: Auth (info), Department (navy), Employee (teal), Session (warning/netral).
- Jangan menampilkan metadata mentah di tabel.

### 18B.2 Audit Detail

Gunakan dialog/drawer:

- Header: action + waktu.
- Ringkasan aktor, entitas, IP, user agent, request ID (jika ada).
- Section “Perubahan”:
  - Daftar field yang berubah.
  - Before → after dalam dua kolom.
  - Nilai null ditampilkan “—”.
  - Field sensitif disamarkan.
- Metadata JSON mentah tidak ditampilkan kecuali dalam mode debug yang diotorisasi (tidak untuk MVP).
- Tombol tutup; tidak ada edit.

### 18B.3 States

- Loading skeleton.
- Empty: “Belum ada aktivitas.”
- No results: reset filter.
- Error: retry.
- Metadata besar: scroll di dalam dialog, bukan memenuhi halaman.

---

## 19. Table Component Guidelines

### 19.1 Anatomi

```text
┌──────────────────────────────────────────────┐
│ Toolbar: search, filter, bulk actions        │
├──────────────────────────────────────────────┤
│ Header: sortable columns, checkbox           │
├──────────────────────────────────────────────┤
│ Row 1                                        │
│ Row 2                                        │
│ ...                                          │
├──────────────────────────────────────────────┤
│ Footer: result count, pagination             │
└──────────────────────────────────────────────┘
```

### 19.2 Spesifikasi

- Container: desktop putih, border `#E8ECF2`, radius 12px, overflow hidden.
  Di mobile container section transparan (tanpa background/border/shadow);
  hanya card item dan input search yang memakai background putih.
- Desktop grid: mengisi sisa tinggi viewport dengan padding bawah 20px,
  `overflow-auto`, scrollbar tipis (`.data-scroll`), header sticky. Tidak ada
  tinggi tetap 520px.
- Kontrak tinggi: Shell luar `min-h-dvh lg:h-dvh lg:overflow-hidden`; Shell
  dalam `flex min-h-dvh flex-col lg:h-dvh lg:min-h-0 lg:pl-[250px]`; topbar
  `h-16 shrink-0`; `main` `flex min-h-0 flex-1 flex-col ... lg:pb-5`; container
  `flex min-h-0 flex-1 flex-col`; root halaman `lg:min-h-0 lg:flex-1`; section
  `lg:min-h-0 lg:flex-1 lg:overflow-hidden`; area tabel `lg:min-h-0 lg:flex-1
  overflow-auto`; toolbar dan pagination `shrink-0`.
- Invariant desktop: `documentElement.scrollHeight === window.innerHeight`;
  bottom section = `window.innerHeight - 20`; hanya area tabel yang scroll.
  Mobile tidak memakai `h-dvh`/`overflow-hidden`, sehingga scroll dokumen tetap
  normal.
- Jangan menghapus `min-h-0` pada salah satu level flex; tanpa itu flex item
  memakai `min-height: auto` dan menolak menyusut, sehingga grid melebihi
  viewport.
- Table: `table-fixed`; lebar kolom ditentukan `<colgroup>` dengan persentase
  yang totalnya 100%.
- Header: background `surface-secondary` (`#F7F9FC`), teks 12px uppercase,
  weight 500, `#667085`.
- Row: padding vertikal 16px, border-top 1px.
- Hover: `surface-secondary/60`.
- Cell padding: 16px horizontal (aksi 12px).
- Teks angka: `tabular-nums`.
- Aksi: tombol ikon tiga titik 32–36px dengan `aria-label`.
- Scroll horizontal tidak muncul karena `table-fixed` dan lebar kolom
  persentase; konten panjang di-truncate dengan `title` untuk teks lengkap.

### 19.3 Sorting

- Klik header mengubah `asc` → `desc` → default.
- Indicator panah hanya pada kolom aktif.
- `aria-sort` di `<th>`.
- Jangan sorting client-side untuk data terpaginasi server.

### 19.4 Pagination

- Desktop: tampilkan rentang dan total (“{from}-{to} dari {total} data”),
  tombol “Sebelumnya”/“Berikutnya”, dan indikator “{page} / {pages}”.
- Tombol disabled pada batas dan saat `busy`/loading.
- Page size tetap 10; belum ada pemilih ukuran halaman.
- Saat pindah halaman/filter, scroll container tabel kembali ke atas.
- Mobile: tidak memakai pagination; memakai infinite scroll 10 item per
  halaman dengan sentinel `IntersectionObserver`.

### 19.5 Row Actions

- Gunakan menu tiga titik (`MoreHorizontal`) untuk aksi baris.
- Menu dirender via portal, diposisikan `fixed`; menutup saat scroll, resize,
  Escape, atau klik overlay.
- Aksi destruktif dipisahkan dengan divider dan warna danger.
- Aksi disabled harus menjelaskan alasan via `title`/tooltip:
  - hapus departemen: “Pindahkan atau hapus karyawan terlebih dahulu”;
  - hapus karyawan: “Nonaktifkan karyawan terlebih dahulu”.
- Menu memakai `aria-haspopup="menu"`, `aria-expanded`, `role="menu"`, dan
  `role="menuitem"`.
- Jangan hanya mengandalkan icon tanpa label aksesibel.

### 19.6 Bulk Selection — Future

- Checkbox hanya jika ada aksi bulk nyata.
- Header checkbox: select all page, bukan seluruh dataset, kecuali dijelaskan.
- Tampilkan bulk action bar saat ada selection.
- Belum termasuk MVP.

### 19.7 Responsive Table

- Desktop (≥ 1024px): tabel penuh dengan kolom tetap dan scroll internal.
- Tablet (< 1024px): **bukan** tabel horizontal scroll; memakai card list agar
  status, departemen, dan aksi tetap terbaca.
- Mobile: card list + infinite scroll.
- Jangan memotong kolom status/aksi.
- Sediakan `aria-label` pada tabel dan `aria-busy` saat loading.

### 19.8 Table States

| State | Tampilan |
|---|---|
| Loading awal | Skeleton rows sesuai bentuk kolom (8 baris desktop, 3 card mobile). |
| Refetch | Data lama tetap tampil dengan opacity 60%; tidak ada spinner overlay. |
| Empty | Empty state di dalam tabel (colspan penuh) atau container card. |
| Error | `ErrorAlert` + tombol “Coba lagi”. |
| No results | Empty state + tombol “Reset filter”. |
| Load more gagal | Sentinel mobile berubah menjadi tombol retry. |

### 19.9 Implementasi Fixed Grid

- `DataTable` menerima `columns` dengan `width`, `skeleton`, dan `render`.
- Setiap kolom merender skeleton yang menyerupai konten asli.
- Container scroll memakai `scrollRef` agar halaman dapat scroll ke atas saat
  filter/halaman berubah.
- Tabel disembunyikan di bawah `lg`; `MobileList` yang mengambil alih.
- `usePagedList` menjadi satu sumber data untuk kedua mode, sehingga filter dan
  total konsisten.

Rantai class aktual (harus dipertahankan saat refactor):

```text
Shell luar       min-h-dvh lg:h-dvh lg:overflow-hidden
Shell dalam      flex min-h-dvh flex-col lg:h-dvh lg:min-h-0 lg:pl-[250px]
Topbar           h-16 shrink-0
Main             flex min-h-0 flex-1 flex-col lg:px-8 lg:pb-5 lg:pt-8
Container        mx-auto flex min-h-0 w-full max-w-[1280px] flex-1 flex-col
Root halaman     flex flex-col lg:min-h-0 lg:flex-1
Section          mt-6 flex flex-col ... lg:min-h-0 lg:flex-1 lg:overflow-hidden
Toolbar          shrink-0
Area tabel       data-scroll hidden overflow-auto lg:block lg:min-h-0 lg:flex-1
Pagination       shrink-0
```

Alasan teknis:

- `h-dvh` memberi tinggi definite pada rantai flex desktop. Tanpa tinggi
  definite, `flex-1` hanya membagi ruang yang tersedia dan tidak bisa
  mengunci tinggi ke viewport.
- `min-h-0` wajib di setiap flex item yang harus menyusut. Default
  `min-height: auto` membuat item mengikuti min-content (mis. tinggi tabel),
  sehingga halaman memanjang.
- `overflow-hidden` pada Shell desktop mencegah scroll dokumen; scroll hanya
  terjadi di area tabel.
- `pb-5` (20px) pada `main` menghasilkan jarak bawah card tepat 20px dari tepi
  viewport. Jangan menggantinya dengan margin pada section tanpa menghitung
  ulang.
- `shrink-0` pada toolbar/pagination menjamin kontrol tetap terlihat saat
  viewport pendek; area tabel yang dikorbankan lebih dulu.
- Mobile tidak memakai `lg:` sehingga `min-h-dvh` + scroll dokumen normal tetap
  berlaku.

Verifikasi:

1. Desktop ≥ 1024px: `documentElement.scrollHeight === window.innerHeight`.
2. `window.innerHeight - section.getBoundingClientRect().bottom === 20`.
3. Perkecil tinggi viewport: tinggi area tabel berkurang, pagination tetap
   terlihat, tidak ada scroll dokumen.
4. Mobile < 1024px: halaman boleh scroll; card list dan infinite scroll tetap
   berfungsi.
5. `npm run build` lulus.

Hasil verifikasi headless terakhir: viewport dalam 804px → `scrollHeight` 804,
`clientHeight` 804, gap 20px, tinggi tabel 466px; viewport 604px → 604/604,
gap 20px, tinggi tabel 266px.

---

## 20. Form Component Guidelines

### 20.1 Struktur Field

```text
Label *                      (helper)
[Input]
Error message
Helper text
```

- Label selalu di atas, terhubung via `htmlFor`/`id`.
- Required indicator `*` dengan `aria-hidden` dan `aria-required` pada input.
- Error: border danger + pesan 13px danger + `aria-describedby`.
- Helper: 12–13px muted.
- Placeholder bukan pengganti label.

### 20.2 Ukuran

| Komponen | Tinggi |
|---|---:|
| Input/select | 42px |
| Textarea | min 96px |
| Button | 40px |
| Small button | 32px |

### 20.3 Validasi

- Validasi saat blur dan submit; jangan memvalidasi setiap keystroke hingga mengganggu.
- Setelah field pernah error, revalidasi saat berubah.
- Pesan spesifik dan actionable.
- Server error dipetakan ke field.
- Form submit disabled saat pending; jangan double-submit.

### 20.4 Field Types

- **Text:** normalisasi trim saat submit; tampilkan uppercase untuk kode jika diinginkan.
- **Email:** `type="email"`, `inputMode="email"`, autocomplete.
- **Phone:** `type="tel"`, `inputMode="tel"`; jangan memaksa format saat mengetik.
- **Select:** placeholder “Pilih…”; opsi aktif; disabled state jelas. Jangan
  memakai `<select>` native. Gunakan komponen `Select` reusable dengan trigger
  class `.field` + `.select-trigger` (tinggi 42px, border/radius/focus sama
  dengan input teks) dan chevron SVG 16px yang berputar saat terbuka.
  Panel dropdown dirender lewat portal ke `document.body` agar tidak terpotong
  modal: `role="listbox"`, opsi `role="option"`, `aria-selected`, background
  putih, border `#E8ECF2`, radius 8px, `shadow-card`, tinggi maksimum 280px,
  scroll tipis. Lebar panel sama dengan lebar trigger; gunakan `box-border` agar
  scrollbar vertikal ikut diperhitungkan di dalam lebar panel. Opsi terpilih
  memakai tint `#EEF3FA`, teks primary, dan ikon centang. Posisi panel di bawah
  trigger; jika ruang bawah tidak cukup, buka ke atas dengan tinggi yang sudah
  dikurangi gap 6px dan margin tepi 8px. Scroll di dalam panel tidak menutup
  dropdown (`overscroll-contain`); scroll di luar panel atau resize menutup.
  Keyboard: Enter/Space/ArrowDown/ArrowUp membuka, ArrowUp/ArrowDown/Home/End
  navigasi, Enter/Space memilih, Escape/Tab menutup, fokus tetap di trigger
  dengan `aria-activedescendant`. Filter memakai opsi “Semua …”; form memakai
  placeholder “Pilih …”. Opsi disabled dilewati navigasi dan tidak dapat dipilih.
  Select filter dan select form wajib memakai komponen yang sama. Spesifikasi
  lengkap ada di §20.7.
- **Date:** date input; tampilkan format lokal di detail, kirim `YYYY-MM-DD`.
- **Textarea:** counter karakter; resize vertical.
- **Search:** icon, clear button, debounce.

### 20.5 Layout

- Desktop: 2 kolom, gap 16–24px; field panjang full width.
- Mobile: 1 kolom.
- Section dipisahkan heading 15–16px dan spacing 24px.
- Footer form: border-top, tombol kanan, sticky di mobile jika form panjang.

### 20.6 Error Ringkasan

Jika > 3 error:

- Tampilkan alert di atas form: “Periksa 4 field yang ditandai.”
- Fokus ke field error pertama.
- Jangan menghilangkan input pengguna.

### 20.7 Select Component Specification

Spesifikasi ini mengikat implementasi `frontend/src/components/Select.tsx`.
Semua select filter dan form wajib memakai komponen ini; tidak boleh ada
`<select>` native baru.

#### 20.7.1 Anatomi

```text
┌──────────────────────────────────────────────┐
│ Label *                                      │
│ ┌──────────────────────────────────────────┐ │
│ │ Nilai terpilih / placeholder         ⌄   │ │  trigger 42px
│ └──────────────────────────────────────────┘ │
│ ┌──────────────────────────────────────────┐ │
│ │ ✓ Opsi terpilih                          │ │  panel listbox
│ │   Opsi aktif                             │ │  lebar = trigger
│ │   Opsi disabled                          │ │  scroll internal
│ └──────────────────────────────────────────┘ │
│ Error message / helper                       │
└──────────────────────────────────────────────┘
```

- Trigger: `<button type="button" role="combobox">`; bukan `<select>`.
- Panel: `role="listbox"`; setiap item `role="option"`.
- Panel dirender lewat portal ke `document.body`; bukan di dalam stacking
  context modal.
- Overlay penutup transparan `fixed inset-0 z-[80]`; panel `z-[81]`; modal
  `z-[60]`.

#### 20.7.2 Ukuran dan positioning

Pseudocode yang menjadi acuan:

```ts
const rect = trigger.getBoundingClientRect();
const gap = 6;
const edge = 8;
const desired = Math.min(280, options.length * 38 + 8);

const spaceBelow = window.innerHeight - rect.bottom - gap - edge;
const spaceAbove = rect.top - gap - edge;
const openUp = spaceBelow < desired && spaceAbove > spaceBelow;
const available = openUp ? spaceAbove : spaceBelow;
const maxHeight = Math.max(0, Math.min(280, available));

const width = rect.width; // WAJIB sama dengan trigger
const left = Math.max(
  edge,
  Math.min(rect.left, window.innerWidth - width - edge),
);
const top = openUp
  ? Math.max(edge, rect.top - maxHeight - gap)
  : rect.bottom + gap;
```

Aturan:

- Lebar panel **selalu sama dengan lebar trigger**. Jangan memakai minimum width
  yang membuat panel lebih lebar dari input.
- `box-border` wajib agar border, padding, dan scrollbar vertikal dihitung di
  dalam lebar total. Dengan begitu munculnya scrollbar tidak menambah lebar
  panel.
- `overflow-x-hidden` dan `truncate` mencegah scrollbar horizontal dan teks
  panjang memperlebar panel.
- Tinggi maksimum 280px; tinggi aktual mengikuti jumlah opsi dan ruang
  viewport.
- Buka ke atas jika ruang bawah tidak cukup dan ruang atas lebih besar.
  `maxHeight` harus dikurangi gap 6px dan margin tepi 8px agar panel tidak
  melewati viewport.
- Panel memakai `position: fixed`; koordinat dihitung dari viewport. Jangan
  memakai `absolute` di dalam modal karena akan terpotong `overflow`.

#### 20.7.3 Scroll behavior

- Panel: `overflow-y-auto` + class `.data-scroll` (scrollbar tipis) +
  `overscroll-contain`.
- Listener scroll global memakai capture untuk menutup panel saat halaman/modal
  di luar panel di-scroll. Event yang target-nya berada di dalam panel harus
  diabaikan; jika tidak, scroll daftar opsi akan menutup dropdown (bug).
- Sorotan opsi aktif digulir dengan menyesuaikan `panel.scrollTop`; jangan
  memakai `scrollIntoView()` karena dapat menggulir ancestor dan memicu event
  scroll global.
- Panel menutup saat: klik overlay, Escape, Tab, resize, scroll di luar panel,
  atau setelah memilih opsi.
- Scroll di dalam panel tidak boleh menutup panel dan tidak boleh merambat ke
  halaman/modal.

#### 20.7.4 States

| State | Trigger | Panel/opsi |
|---|---|---|
| Default | `.field` putih, border `control-border`, teks ink | Tertutup |
| Hover | Border mengikuti `.field:hover` | — |
| Focus | Border primary + ring primary/10 | — |
| Open | Chevron rotate 180° | Panel muncul, opsi terpilih ditandai |
| Placeholder | `data-placeholder="true"` → teks muted | — |
| Disabled | `.field:disabled`: surface-secondary, muted, `not-allowed` | Opsi disabled opacity 60%, tidak bisa dipilih |
| Selected option | — | Background `#EEF3FA`, teks primary, weight medium, ikon centang |
| Active option | — | Background `surface-secondary` |
| Error | Target: border danger + `aria-invalid` | Saat ini error validasi ditampilkan oleh form di bawah field; komponen belum menerima prop `error` |

Catatan: jika menambahkan state error pada Select, tambahkan prop `error`,
`aria-invalid`, `aria-describedby`, dan border danger yang konsisten dengan
`.field`; jangan membuat style error terpisah di halaman.

#### 20.7.5 Keyboard map

| Key | Tertutup | Terbuka |
|---|---|---|
| Enter | Buka panel | Pilih opsi aktif |
| Space | Buka panel | Pilih opsi aktif |
| ArrowDown | Buka panel | Opsi berikutnya (wrap, skip disabled) |
| ArrowUp | Buka panel | Opsi sebelumnya (wrap, skip disabled) |
| Home | — | Opsi enabled pertama |
| End | — | Opsi enabled terakhir |
| Escape | — | Tutup tanpa memilih |
| Tab | — | Tutup, lanjut fokus |

Fokus tetap berada di trigger selama panel terbuka; opsi aktif ditandai dengan
`aria-activedescendant`.

#### 20.7.6 ARIA

- Trigger:
  - `role="combobox"`
  - `aria-haspopup="listbox"`
  - `aria-expanded="true|false"`
  - `aria-controls="<listbox-id>"` saat terbuka
  - `aria-activedescendant="<option-id>"` saat ada opsi aktif
  - `aria-required="true"` jika required
  - `aria-label` untuk filter yang tidak punya label visual
- Panel:
  - `role="listbox"`
  - `aria-label` dari `ariaLabel`, placeholder, atau “Pilihan”
- Opsi:
  - `role="option"`
  - `aria-selected="true|false"`
  - `aria-disabled="true"` untuk opsi disabled
- ID listbox/opsi memakai `useId()` agar unik dan tidak bentrok saat ada
  beberapa select di halaman/modal.

#### 20.7.7 Integrasi

- `FilterSelect` di `components/ListToolbar.tsx` adalah wrapper tipis:
  meneruskan `options`, `value`, `onChange`, `ariaLabel`, dan `className`.
- Form departemen/karyawan memakai `id` + `<label htmlFor>`.
- Filter memakai option kosong “Semua …”; form departemen memakai placeholder
  “Pilih departemen” dan `required`.
- Nilai departemen memakai `public_id`, bukan ID numerik.
- Opsi disabled harus tetap terlihat dan dilewati navigasi keyboard.

#### 20.7.8 Anti-pattern

- Jangan memakai `<select>` native untuk filter/form.
- Jangan membuat panel lebih lebar dari trigger.
- Jangan memakai `width: max-content` atau minimum width pada panel.
- Jangan menutup panel saat panel sendiri di-scroll.
- Jangan memakai `scrollIntoView()` untuk opsi aktif.
- Jangan menaruh panel di dalam modal tanpa portal; panel akan terpotong.
- Jangan memakai z-index panel di bawah modal.
- Jangan menghapus `box-border`; scrollbar bisa menambah lebar panel.
- Jangan membuat komponen select duplikat per halaman.

#### 20.7.9 QA checklist

- [ ] Lebar panel = lebar trigger di filter dan form, termasuk saat scrollbar
      muncul.
- [ ] Panel tidak terpotong di dalam modal dan berada di atas overlay modal.
- [ ] Scroll daftar opsi panjang tidak menutup panel.
- [ ] Scroll halaman/modal di luar panel menutup panel.
- [ ] Keyboard: Enter/Space/Arrow/Home/End/Escape/Tab berfungsi.
- [ ] Opsi disabled tidak bisa dipilih dan dilewati.
- [ ] Placeholder muted; opsi terpilih punya centang dan tint.
- [ ] Viewport pendek: panel membuka ke atas dan tidak melewati tepi layar.
- [ ] Tidak ada `<select>` native di halaman.
- [ ] `npm run build` lulus.

---

## 21. Dialog & Confirmation Guidelines

### 21.1 Ukuran

| Tipe | Lebar |
|---|---:|
| Confirmation | 400–440px |
| Form dialog | 480–560px |
| Large dialog | 640px |

Mobile: lebar penuh dengan margin 16px; tinggi maksimal viewport dengan body scroll.

### 21.2 Struktur

- Title jelas (15–18px, weight 600).
- Deskripsi singkat.
- Body (opsional).
- Footer: secondary action kiri, primary/danger kanan.
- Close icon kanan atas dengan accessible label.

### 21.3 Perilaku

> **Implementasi saat ini:** Escape, klik overlay, dan tombol close sudah
> berfungsi. Focus trap dan pengembalian fokus ke trigger belum
> diimplementasikan. Scroll body belum dikunci saat dialog terbuka.

- Focus trap; fokus awal ke elemen aman (bukan tombol destruktif). *(target)*
- `Esc` menutup, kecuali proses sedang berjalan.
- Klik overlay menutup; untuk konfirmasi yang sedang diproses, handler menolak
  menutup.
- Kembalikan fokus ke trigger setelah tutup. *(target)*
- Scroll body dikunci saat dialog terbuka. *(target)*
- Tombol destruktif tidak menjadi default focus. *(target)*

### 21.4 Konfirmasi Destruktif

Wajib untuk:

- delete department,
- delete employee,
- deactivate department/employee,
- revoke session,
- logout jika ada pekerjaan belum disimpan (opsional).

Gunakan kata kerja spesifik: “Hapus departemen”, bukan “OK”.

---

## 22. Status Badge System

### 22.1 Prinsip

- Badge menunjukkan status, bukan aksi.
- Gunakan label teks; warna bukan satu-satunya pembeda.
- Ukuran proporsional: padding 4–6px 8–10px, radius 6px, font 12px weight 500.
- Jangan memakai badge untuk setiap metadata; hanya status penting.

### 22.2 Mapping

| Domain | Nilai | Label | Gaya |
|---|---|---|---|
| Department | `active` | Aktif | Success tint |
| Department | `inactive` | Nonaktif | Neutral tint |
| Employee | `active` | Aktif | Success tint |
| Employee | `inactive` | Nonaktif | Neutral tint |
| Session | current | Sesi ini | Info tint + dot |
| Session | active | Aktif | Success tint |
| Session | revoked | Dicabut | Neutral/danger tint |
| Audit | auth | Autentikasi | Info tint |
| Audit | department | Departemen | Navy tint |
| Audit | employee | Karyawan | Teal tint |

### 22.3 Token Warna Badge — [REKOMENDASI]

```css
--badge-success-bg: #ECFDF3;
--badge-success-text: #067647;

--badge-neutral-bg: #F2F4F7;
--badge-neutral-text: #475467;

--badge-warning-bg: #FFFAEB;
--badge-warning-text: #B54708;

--badge-danger-bg: #FEF3F2;
--badge-danger-text: #B42318;

--badge-info-bg: #EFF8FF;
--badge-info-text: #175CD3;
```

Warna di atas adalah tint tambahan untuk badge; tetap merujuk semantic colors existing (`success #16A36A`, `warning #E9A23B`, `danger #DC5252`, `info #3182CE`). Jangan mengganti token existing.

### 22.4 Aturan

- Jangan memakai merah untuk “nonaktif”.
- Jangan memakai hijau untuk “baru dibuat” kecuali benar-benar sukses.
- Badge tidak boleh menjadi tombol.
- Status yang dapat berubah harus memiliki aksi terpisah.

---

## 23. Loading, Empty, and Error States

### 23.1 Loading

- **Initial page:** skeleton layout, bukan spinner penuh.
- **Table:** 5–8 skeleton rows.
- **Cards:** skeleton angka dan label.
- **Button:** spinner 16px + label “Menyimpan…”.
- **Inline:** spinner kecil di samping teks.
- **Refetch:** pertahankan data lama + indikator halus.
- Jangan menampilkan angka 0 sementara data belum ada.
- Skeleton harus mengikuti layout final untuk mengurangi layout shift.

### 23.2 Empty States

Anatomi:

```text
[Icon dalam lingkaran lembut]
Judul singkat
Deskripsi 1–2 baris
[Primary action] (opsional)
```

Aturan:

- Bedakan “belum ada data” vs “tidak ada hasil filter”.
- Gunakan bahasa netral, bukan menyalahkan pengguna.
- Sediakan aksi reset filter untuk no-results.
- Empty state tidak boleh menutupi toolbar yang masih berguna.
- Untuk dashboard, empty state dapat mengarahkan ke langkah pertama (tambah departemen).

Contoh copy:

- “Belum ada departemen. Tambahkan departemen pertama untuk mulai mengelola struktur organisasi.”
- “Tidak ada karyawan yang cocok dengan filter. Coba ubah kata kunci atau reset filter.”
- “Belum ada aktivitas tercatat.”

### 23.3 Error States

- **Field error:** inline di bawah input.
- **Form error:** alert di atas form + field errors.
- **Page error:** card/panel dengan icon, judul “Gagal memuat data”, deskripsi singkat, tombol “Coba lagi”.
- **Partial error:** panel kecil; bagian lain tetap tampil.
- **Network error:** “Tidak dapat terhubung ke server. Periksa koneksi Anda.”
- **401:** redirect login dengan pesan sesi berakhir.
- **403:** halaman akses ditolak, bukan empty state.
- **404:** halaman data tidak ditemukan + kembali.
- Jangan menampilkan stack trace, nama tabel, atau pesan database.

### 23.4 Toast

> **Implementasi saat ini:** tidak memakai toast. Feedback mutasi
> create/update/delete memakai `ResponseModal` (ikon + pesan + tombol Tutup).
> Error validasi form tetap inline di dalam modal form.

- Posisi kanan atas (desktop) / atas (mobile).
- Durasi 4-6 detik; error penting tidak auto-hilang terlalu cepat.
- Maksimal 3 toast terlihat.
- Gunakan untuk feedback mutasi sukses/gagal yang tidak memerlukan tindakan.
- Untuk error validasi form, gunakan inline, bukan toast saja.
- Toast harus dapat ditutup dan accessible (`role="status"`/`alert`).

---

## 24. Component Library Guidelines

### 24.1 Komponen Dasar yang Diperlukan

| Komponen | State wajib |
|---|---|
| `Button` | default, hover, active, focus, disabled, loading |
| `IconButton` | default, hover, focus, disabled; accessible label |
| `Input` | default, hover, focus, error, disabled, readonly |
| `Select` | default, open, focus, error, disabled, placeholder, selected option, active option |
| `Textarea` | default, focus, error, disabled; counter |
| `Checkbox` | checked, unchecked, indeterminate, disabled |
| `Badge` | success, neutral, warning, danger, info |
| `Card` | default, interactive (opsional) |
| `Table` | loading, empty, error, selected, sorted |
| `Pagination` | default, disabled, current |
| `Dialog` | open, loading, destructive |
| `Alert` | info, success, warning, danger |
| `Toast` | success, error, info |
| `Skeleton` | text, card, row |
| `EmptyState` | default, no-results |
| `ErrorState` | page, panel, inline |
| `DropdownMenu` | open, focus, disabled item |
| `Tabs` (future) | active, disabled |
| `Tooltip` | hover/focus; accessible |

### 24.2 Aturan Komponen

- Reusable dan tidak mengandung business logic.
- Props eksplisit; hindari boolean berlebihan.
- Gunakan `className` terbatas; jangan override warna sesuka hati.
- Semua state harus dirancang, termasuk disabled/loading/error.
- Komponen harus mendukung keyboard dan `aria-*` yang relevan.
- Jangan membuat komponen baru jika komponen existing bisa diperluas.
- Jangan menduplikasi style; gunakan token.

### 24.3 Naming

- Komponen: PascalCase (`StatCard`, `DepartmentTable`).
- Props: camelCase.
- CSS token: kebab-case dengan prefix (`--color-primary`).
- File komponen: PascalCase; util: camelCase.
- Hindari nama generik seperti `DataTable2`, `NewModal`.

### 24.4 Icon

- Standar: Lucide Icons (outline), ukuran 18–20px, stroke 1.75–2.
- Library Lucide belum ada di `package.json`. Saat ini beberapa ikon memakai inline SVG.
- Jika Lucide ditambahkan, gunakan secara konsisten; jangan mencampur banyak gaya ikon.
- Icon dekoratif `aria-hidden`; icon-only button wajib `aria-label`.
- Jangan memakai emoji sebagai ikon UI.

---

## 25. Page Templates

### 25.1 List Page Template

```text
Page title + subtitle                         [Primary action]
┌─ Card ──────────────────────────────────────────────────────┐
│ [Search] [Filter 1] [Filter 2] [Reset]                       │
├──────────────────────────────────────────────────────────────┤
│ Desktop ≥1024px: data grid mengisi sisa viewport (20px dari bawah)  │
│ Mobile <1024px: card list + skeleton + infinite sentinel             │
├──────────────────────────────────────────────────────────────┤
│ Desktop: pagination 10/page                                  │
│ Mobile: sentinel "Gulir untuk memuat..."                     │
└──────────────────────────────────────────────────────────────┘
```

- Desktop: toolbar dan daftar berada dalam satu card putih. Mobile: container
  toolbar + daftar transparan; hanya card item dan input search yang putih.
- Search memakai submit; filter select langsung memuat ulang.
- Reset hanya muncul saat ada filter aktif.
- Empty state karena filter menyediakan tombol reset.
- Semua mutasi memakai modal form + modal respons.

### 25.2 Detail Page Template

```text
Breadcrumb
[Title + status]                              [Actions]
[Metadata / summary cards]
[Section 1]
[Section 2]
```

### 25.3 Form Page Template

```text
Breadcrumb
Page title
[Section: Informasi]
[Fields...]
[Section: Status]
[Fields...]
[Footer: Batal | Simpan]
```

### 25.4 Auth Page Template

- Login existing dipertahankan: background `#F7F9FC`, card putih, logo resmi `logo.png`, gradient accent tipis.
- Jangan menambahkan elemen marketing besar.
- Error alert di atas form.
- Link “Hubungi administrator” tetap ada, tampil inline setelah “Belum punya akses?” tanpa icon dan hanya link tersebut yang bold berwarna primary.
- Setelah login sukses, arahkan ke `/departments` (dashboard overview belum ada).

---

## 26. Content & Microcopy

### 26.1 Prinsip

- Bahasa Indonesia, ringkas, profesional.
- Gunakan kata kerja aktif.
- Hindari istilah teknis internal (“payload”, “endpoint”, “token”) di UI pengguna.
- Konsisten: “Departemen”, “Karyawan”, “Nonaktif”, “Sesi”.
- Gunakan “Anda” secara konsisten atau hindari kata ganti.

### 26.2 Contoh Label

| Konteks | Copy |
|---|---|
| Create department | “Tambah Departemen” |
| Save | “Simpan” |
| Cancel | “Batal” |
| Delete | “Hapus” |
| Deactivate | “Nonaktifkan” |
| Reactivate | “Aktifkan Kembali” |
| Search | “Cari…” |
| Empty | “Belum ada data.” |
| Error | “Gagal memuat data.” |
| Retry | “Coba lagi” |
| Success | “Perubahan berhasil disimpan.” |

### 26.3 Konfirmasi

Gunakan nama objek spesifik:

- “Hapus departemen Engineering?”
- “Nonaktifkan karyawan Budi Santoso?”
- “Cabut sesi ini?”

Jangan memakai “Apakah Anda yakin?” tanpa konteks.

---

## 27. Responsive Behavior per Page

### 27.1 Dashboard

- 1440/1280: 4 stat cards, 2 kolom konten.
- 768: 2 stat cards per baris; panel bertumpuk.
- 375: 1–2 stat cards; chart bar tetap terbaca; activity 5 item.

### 27.2 Department List

- Desktop (≥ 1024px): data grid mengisi sisa tinggi viewport dan berhenti 20px
  di atas tepi bawah viewport, kolom tetap, scroll internal, pagination 10 per
  halaman. Kontrak tinggi lengkap ada di §19.9.
- Tablet (< 1024px): card list; toolbar stacked; search full width; select
  status full width; reset full width; container daftar transparan sehingga
  hanya card dan search yang memiliki background.
- Select status memakai komponen §20.7; panel mengikuti lebar trigger dan tidak
  boleh terpotong.
- Mobile: card list + infinite scroll; tombol tambah berada di bawah header.
- Card menampilkan nama, kode, deskripsi (2 baris), status, jumlah karyawan,
  dan menu tiga titik.

### 27.3 Employee List

- Desktop (≥ 1024px): data grid mengisi sisa viewport (20px dari bawah) dengan
  7 kolom tetap, pagination 10 per halaman. Kontrak tinggi lengkap ada di §19.9.
- Tablet/mobile (< 1024px): card list dengan avatar inisial, nama, employee ID,
  status, menu tiga titik, lalu grid 2 kolom berisi email, telepon, departemen,
  jabatan, dan tanggal bergabung. Container daftar transparan; hanya card dan
  search yang memiliki background.
- Filter departemen/status dan select form memakai komponen §20.7; nilai
  departemen tetap `public_id`.
- Infinite scroll 10 item per halaman; skeleton card saat loading.
- Jangan menyembunyikan status dan departemen.

### 27.4 Forms

- Desktop: 2 kolom.
- Tablet: 1–2 kolom sesuai lebar.
- Mobile: 1 kolom, footer sticky.
- Dialog form menjadi full-screen sheet di mobile jika konten panjang.

### 27.5 Sidebar

- Desktop ≥ 1024px: 250px fixed/sticky.
- Tablet: collapsed/icon-only atau drawer.
- Mobile < 768px: drawer + overlay; tombol close; fokus trap.
- Jangan membuat sidebar dark penuh.

---

## 28. Accessibility Checklist

- [ ] Semua input memiliki label yang terhubung.
- [ ] Error field memakai `aria-invalid` dan `aria-describedby`.
- [ ] Tombol icon-only memiliki `aria-label`.
- [ ] Dialog: `role="dialog"`, `aria-modal="true"`, labelledby, focus trap, Esc, return focus.
- [ ] Tabel: caption, `th scope`, `aria-sort`.
- [ ] Status tidak hanya warna; ada teks.
- [ ] Kontras teks minimal 4.5:1 untuk body, 3:1 untuk teks besar.
- [ ] Focus visible jelas; jangan menghapus outline tanpa pengganti.
- [ ] Navigasi keyboard untuk sidebar, tabel, pagination, dropdown.
- [ ] Select memakai pola combobox/listbox: `role="combobox"`,
      `aria-expanded`, `aria-controls`, `aria-activedescendant`,
      `role="option"`, `aria-selected`, `aria-disabled`; fokus tetap di trigger.
- [ ] Scroll di dalam panel select tidak menutup panel; Escape/Tab menutup.
- [ ] Skip link ke konten utama (rekomendasi).
- [ ] Toast/alert diumumkan (`role="status"`/`alert`).
- [ ] Bahasa halaman `lang="id"` (existing).
- [ ] Reduced motion dihormati untuk animasi.
- [ ] Tidak ada auto-play/animasi mengganggu.
- [ ] Ukuran target sentuh minimal 44×44px pada mobile.

---

## 29. Design Tokens Implementation

Token existing di `styles.css` harus dipertahankan. Tambahan token badge/status di bagian 22 bersifat rekomendasi dan harus ditambahkan ke theme, bukan hardcoded di komponen.

Aturan:

- Gunakan `--color-primary`, `--color-teal`, `--color-border`, dll.
- Jangan menulis hex berulang di JSX.
- Shadow hanya `--shadow-card`/`--shadow-button` yang ada; jangan membuat shadow baru tanpa alasan.
- Radius: control 8px, card 12px.
- Spacing kelipatan 4px.
- Typography Inter; weight 400/500/600/700 saja.

Jika token baru diperlukan:

1. Dokumentasikan di `design.md`.
2. Tambahkan ke `@theme`.
3. Gunakan konsisten.
4. Pastikan kontras.

---

## 30. Design QA Checklist

Sebelum sebuah halaman dianggap selesai:

- [ ] Mengikuti layout sidebar + top navigation existing.
- [ ] Menggunakan token warna/typography/spacing existing.
- [ ] Page title, subtitle, dan primary action konsisten.
- [ ] Tidak ada gradient/shadow/radius berlebihan.
- [ ] Tabel memiliki loading, empty, error, dan no-results state.
- [ ] Form memiliki label, required indicator, helper, dan error state.
- [ ] Dialog destruktif memakai konfirmasi dan copy spesifik.
- [ ] Status memakai badge yang benar dan tidak hanya warna.
- [ ] Responsive diuji pada 1440, 1280, 768, 375px.
- [ ] Desktop grid: tidak ada scroll dokumen, card berhenti 20px di atas
      viewport, hanya area tabel yang scroll, pagination tetap terlihat.
- [ ] Select: lebar panel = lebar trigger, scrollbar dihitung di dalam lebar,
      scroll opsi tidak menutup panel, panel tidak terpotong modal, keyboard dan
      ARIA berfungsi.
- [ ] Keyboard navigation dan focus state berfungsi.
- [ ] Kontras dan label aksesibilitas diperiksa.
- [ ] Tidak ada data placeholder yang terlihat seperti data nyata.
- [ ] Tidak ada teks Inggris yang tidak perlu di UI.
- [ ] Tidak ada komponen duplikat yang bisa memakai komponen existing.
- [ ] Perubahan visual tidak melanggar `design.md` dan tidak menambah design system baru.

---

## 31. Traceability: PRD → Design

| PRD | Design Section |
|---|---|
| AUTH-001, AUTH-002, AUTH-014 | Login template, form, error state (25.4, 20, 23) |
| AUTH-003, AUTH-009, AUTH-010 | Session Management UI (18A) |
| DASH-001–DASH-010 | Dashboard Overview (16) |
| DEPT-001–DEPT-015 | Department Management UI (17) |
| EMP-001–EMP-017 | Employee Management UI (18) |
| AUD-001–AUD-010 | Activity panel (16.5), Audit Log UI (18B), empty/error states |
| NFR-005, NFR-006 | Accessibility (28), Responsive (27) |
| NFR-002 | Loading states (23.1) |

---

## 32. Final Note

Dokumen ini merupakan perluasan dari design system existing. Seluruh warna, tipografi, spacing, layout, sidebar, top navigation, komponen, dan prinsip visual yang sudah ditentukan di bagian 1–13 **dipertahankan**. Penambahan bagian 14–32 hanya melengkapi kebutuhan dashboard tanpa membuat design system baru atau bertentangan dengan identitas Amerta Technology.