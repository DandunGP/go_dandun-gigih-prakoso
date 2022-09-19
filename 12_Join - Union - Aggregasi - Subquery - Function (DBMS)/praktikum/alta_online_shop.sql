-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 19 Sep 2022 pada 12.00
-- Versi server: 10.4.22-MariaDB
-- Versi PHP: 8.1.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `alta_online_shop`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `alamat`
--

CREATE TABLE `alamat` (
  `id_alamat` int(11) NOT NULL,
  `address` varchar(150) NOT NULL,
  `id_userFK` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `operators`
--

CREATE TABLE `operators` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `operators`
--

INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Telkomsel', '2022-09-18 17:44:15', '2022-09-18 17:44:15'),
(2, 'Indosat', '2022-09-18 17:44:15', '2022-09-18 17:44:15'),
(3, 'XL', '2022-09-18 17:44:15', '2022-09-18 17:44:15'),
(4, 'Smartfren', '2022-09-18 17:44:15', '2022-09-18 17:44:15'),
(5, 'Axis', '2022-09-18 17:44:15', '2022-09-18 17:44:15');

-- --------------------------------------------------------

--
-- Struktur dari tabel `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 'Transfer Bank', 1, '2022-09-18 18:06:18', '2022-09-18 18:06:18'),
(2, 'ShopeePay', 1, '2022-09-18 18:06:18', '2022-09-18 18:06:18'),
(3, 'GoPay', 1, '2022-09-18 18:06:18', '2022-09-18 18:06:18');

-- --------------------------------------------------------

--
-- Struktur dari tabel `payment_method_description`
--

CREATE TABLE `payment_method_description` (
  `id_paydes` int(11) NOT NULL,
  `description` varchar(150) DEFAULT NULL,
  `id_paymentFK` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Struktur dari tabel `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `product_type_id` int(11) DEFAULT NULL,
  `operator_id` int(11) DEFAULT NULL,
  `code` varchar(50) NOT NULL,
  `name` varchar(100) NOT NULL,
  `status` smallint(6) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `products`
--

INSERT INTO `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES
(3, 2, 1, 'PKT15GB', 'Paket Data 15GB', 1, '2022-09-18 17:59:11', '2022-09-18 17:59:11'),
(4, 2, 1, 'PKT25GB', 'Paket Data 25GB', 1, '2022-09-18 17:59:11', '2022-09-18 17:59:11'),
(5, 2, 1, 'PKT60GB', 'Paket Data 60GB', 0, '2022-09-18 17:59:11', '2022-09-18 17:59:11'),
(6, 3, 4, 'VCR10GB', 'Voucher 10GB', 1, '2022-09-18 17:59:11', '2022-09-18 17:59:11'),
(7, 3, 4, 'VCR15GB', 'Voucher 15GB', 1, '2022-09-18 17:59:11', '2022-09-18 17:59:11'),
(8, 3, 4, 'VCR20GB', 'Voucher 20GB', 0, '2022-09-18 17:59:11', '2022-09-18 17:59:11');

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_descriptions`
--

CREATE TABLE `product_descriptions` (
  `id` int(11) NOT NULL,
  `description` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `product_descriptions`
--

INSERT INTO `product_descriptions` (`id`, `description`, `created_at`, `updated_at`) VALUES
(1, 'Mengisi Pulsa 15.000', '2022-09-18 18:03:51', '2022-09-18 18:03:51'),
(2, 'Mengisi Pulsa 25.000', '2022-09-18 18:03:51', '2022-09-18 18:03:51'),
(3, '10GB Paket Data + 5GB Paket Malam', '2022-09-18 18:03:51', '2022-09-18 18:03:51'),
(4, '15GB Paket Data + 10GB Paket Malam', '2022-09-18 18:03:51', '2022-09-18 18:03:51'),
(5, '60GB Paket Data', '2022-09-18 18:03:51', '2022-09-18 18:03:51'),
(6, 'Voucher Paket Data 10GB', '2022-09-18 18:03:51', '2022-09-18 18:03:51'),
(7, 'Voucher Paket Data 15GB', '2022-09-18 18:03:51', '2022-09-18 18:03:51'),
(8, 'Voucher Paket Data 20GB', '2022-09-18 18:03:51', '2022-09-18 18:03:51');

-- --------------------------------------------------------

--
-- Struktur dari tabel `product_types`
--

CREATE TABLE `product_types` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `product_types`
--

INSERT INTO `product_types` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'Pulsa', '2022-09-18 17:47:55', '2022-09-18 17:47:55'),
(2, 'Paket Data', '2022-09-18 17:47:55', '2022-09-18 17:47:55'),
(3, 'Voucher', '2022-09-18 17:47:55', '2022-09-18 17:47:55');

-- --------------------------------------------------------

--
-- Struktur dari tabel `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `payment_method_id` int(11) DEFAULT NULL,
  `status` varchar(10) NOT NULL,
  `total_qty` int(11) NOT NULL,
  `total_price` decimal(25,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(1, 2, 2, '1', -1, '30000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(4, 1, 1, '1', 1, '40000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(5, 1, 2, '1', 4, '75000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(6, 1, 3, '1', 2, '20000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(7, 3, 1, '0', 1, '15000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(8, 3, 2, '1', 2, '40000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(9, 3, 3, '1', -2, '20000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(10, 4, 1, '1', 1, '15000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(11, 4, 2, '0', -1, '30000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(12, 4, 3, '1', -2, '20000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(13, 5, 1, '1', 2, '30000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(14, 5, 2, '2', 0, '50000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57'),
(15, 5, 3, '2', 4, '75000.00', '2022-09-19 03:58:57', '2022-09-19 03:58:57');

--
-- Trigger `transactions`
--
DELIMITER $$
CREATE TRIGGER `delete_data_transaction` AFTER DELETE ON `transactions` FOR EACH ROW BEGIN
DECLARE v_transaction_id INT;
SET v_transaction_id = OLD.id;
DELETE FROM transaction_details WHERE transaction_id = v_transaction_id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Struktur dari tabel `transaction_details`
--

CREATE TABLE `transaction_details` (
  `transaction_id` int(11) DEFAULT NULL,
  `product_id` int(11) DEFAULT NULL,
  `status` varchar(10) NOT NULL,
  `qty` int(11) NOT NULL,
  `price` decimal(25,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `transaction_details`
--

INSERT INTO `transaction_details` (`transaction_id`, `product_id`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES
(4, 4, '1', 1, '40000.00', '2022-09-19 04:09:35', '2022-09-19 04:09:35'),
(5, 1, '1', 3, '75000.00', '2022-09-19 04:09:35', '2022-09-19 04:09:35'),
(6, 6, '1', 2, '20000.00', '2022-09-19 04:09:35', '2022-09-19 04:09:35'),
(7, 1, '1', 3, '15000.00', '2022-09-19 04:09:35', '2022-09-19 04:09:35'),
(8, 8, '1', 2, '40000.00', '2022-09-19 04:09:35', '2022-09-19 04:09:35'),
(10, 1, '1', 3, '15000.00', '2022-09-19 04:09:35', '2022-09-19 04:09:35'),
(13, 7, '1', 2, '30000.00', '2022-09-19 04:09:35', '2022-09-19 04:09:35'),
(15, 1, '1', 3, '75000.00', '2022-09-19 04:09:35', '2022-09-19 04:09:35');

--
-- Trigger `transaction_details`
--
DELIMITER $$
CREATE TRIGGER `update_qty_transaction` AFTER DELETE ON `transaction_details` FOR EACH ROW BEGIN
DECLARE v_transaction_id INT;
DECLARE v_transaction_qty INT;
SET v_transaction_id = OLD.transaction_id;
SET v_transaction_qty = OLD.qty;
UPDATE transactions SET total_qty = total_qty - v_transaction_qty WHERE id = v_transaction_id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Struktur dari tabel `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `status` smallint(6) NOT NULL,
  `dob` date NOT NULL,
  `gender` char(1) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `users`
--

INSERT INTO `users` (`id`, `status`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
(1, 1, '2002-11-02', 'M', '2022-09-18 18:10:16', '2022-09-18 18:10:16'),
(2, 1, '1999-07-11', 'M', '2022-09-18 18:10:16', '2022-09-18 18:10:16'),
(3, 1, '1996-02-12', 'M', '2022-09-18 18:10:16', '2022-09-18 18:10:16'),
(4, 1, '2005-05-21', 'F', '2022-09-18 18:10:16', '2022-09-18 18:10:16'),
(5, 1, '2006-02-25', 'F', '2022-09-18 18:10:16', '2022-09-18 18:10:16');

-- --------------------------------------------------------

--
-- Struktur dari tabel `user_payment_method_detail`
--

CREATE TABLE `user_payment_method_detail` (
  `id_uspmd` int(11) NOT NULL,
  `id_userFK` int(11) DEFAULT NULL,
  `id_paymentFK` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `alamat`
--
ALTER TABLE `alamat`
  ADD PRIMARY KEY (`id_alamat`);

--
-- Indeks untuk tabel `operators`
--
ALTER TABLE `operators`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `payment_method_description`
--
ALTER TABLE `payment_method_description`
  ADD PRIMARY KEY (`id_paydes`);

--
-- Indeks untuk tabel `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `product_descriptions`
--
ALTER TABLE `product_descriptions`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `product_types`
--
ALTER TABLE `product_types`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `user_payment_method_detail`
--
ALTER TABLE `user_payment_method_detail`
  ADD PRIMARY KEY (`id_uspmd`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
