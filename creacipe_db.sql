-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jun 14, 2024 at 05:57 AM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.2.4

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
-- Table structure for table `bukuresep`
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
-- Dumping data for table `bukuresep`
--

INSERT INTO `bukuresep` (`id`, `judul`, `bahan`, `langkah`, `created_at`, `updated_at`) VALUES
(1, 'Ayam bakar madu', 'ayam\r\nmadu', 'potong ayam\r\n', '2024-06-13 12:59:36', '2024-06-13 12:59:36'),
(2, 'ayam lada hitam', 'lada hitam , ayam', 'di potong potong', '2024-06-13 12:05:29', '2024-06-13 12:05:29'),
(4, 'astaga', 'gini ', 'rawrs', '2024-06-13 14:21:24', '2024-06-13 14:46:48');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `bukuresep`
--
ALTER TABLE `bukuresep`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `bukuresep`
--
ALTER TABLE `bukuresep`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
