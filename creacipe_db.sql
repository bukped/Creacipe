-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 03 Jul 2024 pada 14.23
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
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data untuk tabel `bukuresep`
--

INSERT INTO `bukuresep` (`id`, `judul`, `bahan`, `langkah`, `created_at`, `updated_at`) VALUES
(1, 'Ayam bakar madu', '500 gram ayam\r\n1 buah jeruk nipis\r\n2 lembar daun jeruk\r\n2 lembar daun salam\r\n2 sdm saus tiram\r\n3 sdm madu', '1. Potong ayam menjadi beberapa bagian. Baluri dengan jeruk nipis, biarkan beberapa menit lalu cuci.\r\n2. Panaskan minyak/margarin, tumis bumbu halus. Masukkan daun jeruk dan daun salam, tumis hingga harum.\r\n3. Masukkan ayam, aduk-aduk.\r\n4. Tambahkan air, ', '2024-06-13 12:59:36', '2024-06-14 07:13:25');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `bukuresep`
--
ALTER TABLE `bukuresep`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `bukuresep`
--
ALTER TABLE `bukuresep`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
