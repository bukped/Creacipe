-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 20 Jul 2024 pada 08.51
-- Versi server: 10.4.32-MariaDB
-- Versi PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `creacipe_db`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `bukuresep`
--

CREATE TABLE `bukuresep` (
  `id` int(11) NOT NULL,
  `judul` varchar(50) NOT NULL,
  `bahan` varchar(255) NOT NULL,
  `langkah` varchar(255) NOT NULL,
  `gambar` varchar(70) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `jenis_masakan_id` int(11) NOT NULL,
  `tingkat_kesulitan_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `bukuresep`
--

INSERT INTO `bukuresep` (`id`, `judul`, `bahan`, `langkah`, `gambar`, `created_at`, `updated_at`, `jenis_masakan_id`, `tingkat_kesulitan_id`) VALUES
(35, 'ayam bakar', '1 kg ayam\r\n200ml kecap manis', '1. rebus ayam\r\n2. campurkan dengan kecap', 'Cara Membuat Ayam Bakar Betutu Bali dan Resep.jpg', '2024-07-20 04:55:03', '2024-07-20 04:55:03', 1, 2),
(36, 'Roti bakar', 'roti tawar\r\nselai rasa\r\nmentega', 'olesi roti tawar dengan mentega terlebih dahulu \r\nlalu panggang dengan api kecil', 'cropped-Roti-Bakar.jpg', '2024-07-20 05:00:34', '2024-07-20 05:00:34', 3, 1);

-- --------------------------------------------------------

--
-- Struktur dari tabel `jenis_masakan`
--

CREATE TABLE `jenis_masakan` (
  `id` int(11) NOT NULL,
  `nama_jenis` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `jenis_masakan`
--

INSERT INTO `jenis_masakan` (`id`, `nama_jenis`) VALUES
(1, 'Masakan Nusantara'),
(2, 'Masakan Barat'),
(3, 'Masakan Asia');

-- --------------------------------------------------------

--
-- Struktur dari tabel `tingkat_kesulitan`
--

CREATE TABLE `tingkat_kesulitan` (
  `id` int(11) NOT NULL,
  `nama_level` varchar(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `tingkat_kesulitan`
--

INSERT INTO `tingkat_kesulitan` (`id`, `nama_level`) VALUES
(1, 'Mudah'),
(2, 'Sedang'),
(3, 'Sulit');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `bukuresep`
--
ALTER TABLE `bukuresep`
  ADD PRIMARY KEY (`id`),
  ADD KEY `jenis_masakan_id` (`jenis_masakan_id`),
  ADD KEY `tingkat_kesulitan_id` (`tingkat_kesulitan_id`);

--
-- Indeks untuk tabel `jenis_masakan`
--
ALTER TABLE `jenis_masakan`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `tingkat_kesulitan`
--
ALTER TABLE `tingkat_kesulitan`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `bukuresep`
--
ALTER TABLE `bukuresep`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=37;

--
-- AUTO_INCREMENT untuk tabel `jenis_masakan`
--
ALTER TABLE `jenis_masakan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT untuk tabel `tingkat_kesulitan`
--
ALTER TABLE `tingkat_kesulitan`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Ketidakleluasaan untuk tabel pelimpahan (Dumped Tables)
--

--
-- Ketidakleluasaan untuk tabel `bukuresep`
--
ALTER TABLE `bukuresep`
  ADD CONSTRAINT `bukuresep_ibfk_2` FOREIGN KEY (`tingkat_kesulitan_id`) REFERENCES `tingkat_kesulitan` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `bukuresep_ibfk_3` FOREIGN KEY (`jenis_masakan_id`) REFERENCES `jenis_masakan` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
