-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 10, 2023 at 05:45 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.0.25

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `latihan_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `operators`
--

CREATE TABLE `operators` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `operators`
--

INSERT INTO `operators` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'arya', '2023-09-09 14:08:45', '2023-09-09 14:08:45'),
(2, 'el', '2023-09-09 14:10:25', '2023-09-09 14:10:25'),
(3, 'prab', '2023-09-09 14:10:25', '2023-09-09 14:10:25'),
(4, 'no', '2023-09-09 14:10:25', '2023-09-09 14:10:25'),
(5, 'ni', '2023-09-09 14:10:25', '2023-09-09 14:10:25');

-- --------------------------------------------------------

--
-- Table structure for table `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `status`, `created_at`, `updated_at`) VALUES
(1, 'gopay', NULL, '2023-09-09 14:37:39', '2023-09-09 14:37:39'),
(2, 'bca', NULL, '2023-09-09 14:37:39', '2023-09-09 14:37:39'),
(3, 'cash', NULL, '2023-09-09 14:37:39', '2023-09-09 14:37:39');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(11) NOT NULL,
  `product_type_id` int(11) DEFAULT NULL,
  `operator_id` int(11) DEFAULT NULL,
  `code` varchar(50) DEFAULT NULL,
  `name` varchar(100) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `product_type_id`, `operator_id`, `code`, `name`, `status`, `created_at`, `updated_at`) VALUES
(3, 2, 1, 'rawr', 'chitato', NULL, '2023-09-09 14:22:19', '2023-09-09 14:22:19'),
(4, 2, 1, 'roar', 'potabee', NULL, '2023-09-09 14:22:19', '2023-09-09 14:22:19'),
(5, 2, 1, 'crunch', 'lays', NULL, '2023-09-09 14:23:40', '2023-09-09 14:23:40'),
(6, 3, 4, 'sweat', 'pocari', NULL, '2023-09-09 14:25:17', '2023-09-09 14:25:17'),
(7, 3, 4, 'diabetes', 'redbull', NULL, '2023-09-09 14:25:17', '2023-09-09 14:25:17'),
(8, 3, 4, 'adem', 'lasegar', NULL, '2023-09-09 14:25:17', '2023-09-09 14:25:17');

-- --------------------------------------------------------

--
-- Table structure for table `product_description`
--

CREATE TABLE `product_description` (
  `id` int(11) NOT NULL,
  `product_id` int(11) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product_description`
--

INSERT INTO `product_description` (`id`, `product_id`, `description`, `created_at`, `updated_at`) VALUES
(1, 1, 'daging pilihan terbaik dari australia', '2023-09-09 14:34:39', '2023-09-09 14:34:39'),
(2, 2, 'ikan terbaik ajalah pokoknya', '2023-09-09 14:34:39', '2023-09-09 14:34:39'),
(3, 3, 'snack kentang', '2023-09-09 14:34:39', '2023-09-09 14:34:39'),
(4, 4, 'snack kentang juga tapi beda tekstur', '2023-09-09 14:34:39', '2023-09-09 14:34:39'),
(5, 5, 'ini juga snack kentang tapi paling enak', '2023-09-09 14:34:39', '2023-09-09 14:34:39'),
(6, 6, 'minuman isotonik', '2023-09-09 14:34:39', '2023-09-09 14:34:39'),
(7, 7, 'energy drink bikin diabetes', '2023-09-09 14:34:39', '2023-09-09 14:34:39'),
(8, 8, 'bikin ilang panas dalam', '2023-09-09 14:34:39', '2023-09-09 14:34:39');

-- --------------------------------------------------------

--
-- Table structure for table `product_types`
--

CREATE TABLE `product_types` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product_types`
--

INSERT INTO `product_types` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'makanan', '2023-09-09 14:13:49', '2023-09-09 14:13:49'),
(2, 'camilan', '2023-09-09 14:13:49', '2023-09-09 14:13:49'),
(3, 'minuman', '2023-09-09 14:13:49', '2023-09-09 14:13:49');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) DEFAULT NULL,
  `payment_method_id` int(11) DEFAULT NULL,
  `status` varchar(10) DEFAULT NULL,
  `total_qty` int(11) DEFAULT NULL,
  `total_price` decimal(25,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `payment_method_id`, `status`, `total_qty`, `total_price`, `created_at`, `updated_at`) VALUES
(1, 1, 1, NULL, 2, '2000.00', '2023-09-09 15:06:38', '2023-09-09 15:17:49'),
(2, 1, 1, NULL, 3, '3000.00', '2023-09-09 15:06:38', '2023-09-09 15:06:38'),
(3, 1, 1, NULL, 4, '4000.00', '2023-09-09 15:06:38', '2023-09-09 15:06:38'),
(4, 2, 2, NULL, 1, '1000.00', '2023-09-09 15:07:20', '2023-09-09 15:07:20'),
(5, 2, 2, NULL, 5, '5000.00', '2023-09-09 15:07:20', '2023-09-09 15:07:20'),
(6, 2, 2, NULL, 4, '4000.00', '2023-09-09 15:07:20', '2023-09-09 15:09:01'),
(7, 3, 3, NULL, 1, '1000.00', '2023-09-09 15:08:06', '2023-09-09 15:09:12'),
(8, 3, 2, NULL, 5, '5000.00', '2023-09-09 15:08:06', '2023-09-09 15:09:22'),
(9, 3, 1, NULL, 3, '3000.00', '2023-09-09 15:08:06', '2023-09-09 15:09:30'),
(10, 4, 2, NULL, 5, '5000.00', '2023-09-09 15:10:00', '2023-09-09 15:10:00'),
(11, 4, 3, NULL, 1, '1000.00', '2023-09-09 15:10:00', '2023-09-09 15:10:00'),
(12, 4, 3, NULL, 3, '3000.00', '2023-09-09 15:10:00', '2023-09-09 15:10:00'),
(13, 5, 2, NULL, 5, '5000.00', '2023-09-09 15:10:25', '2023-09-09 15:10:25'),
(14, 5, 3, NULL, 1, '1000.00', '2023-09-09 15:10:25', '2023-09-09 15:10:25'),
(15, 5, 1, NULL, 1, '1000.00', '2023-09-09 15:10:25', '2023-09-09 15:10:25');

--
-- Triggers `transactions`
--
DELIMITER $$
CREATE TRIGGER `after_transaction_delete` AFTER DELETE ON `transactions` FOR EACH ROW BEGIN
    DELETE FROM transaction_details
    WHERE transaction_id = OLD.id;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `transaction_id` int(11) DEFAULT NULL,
  `product_id` int(11) NOT NULL,
  `status` varchar(10) DEFAULT NULL,
  `qty` int(11) DEFAULT NULL,
  `price` decimal(25,2) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transaction_details`
--

INSERT INTO `transaction_details` (`transaction_id`, `product_id`, `status`, `qty`, `price`, `created_at`, `updated_at`) VALUES
(1, 7, NULL, 4, '1000.00', '2023-09-09 15:33:37', '2023-09-09 15:33:37'),
(1, 4, NULL, 2, '1000.00', '2023-09-09 15:33:37', '2023-09-09 15:33:37'),
(1, 2, NULL, 1, '1000.00', '2023-09-09 15:33:37', '2023-09-09 15:33:37'),
(2, 5, NULL, 1, '1000.00', '2023-09-09 15:34:16', '2023-09-09 15:34:16'),
(2, 4, NULL, 1, '1000.00', '2023-09-09 15:34:16', '2023-09-09 15:34:16'),
(2, 3, NULL, 1, '1000.00', '2023-09-09 15:34:16', '2023-09-09 15:34:16'),
(3, 5, NULL, 3, '1000.00', '2023-09-09 15:34:33', '2023-09-09 15:34:33'),
(3, 5, NULL, 3, '1000.00', '2023-09-09 15:34:33', '2023-09-09 15:34:33'),
(3, 5, NULL, 3, '1000.00', '2023-09-09 15:34:33', '2023-09-09 15:34:33'),
(4, 8, NULL, 1, '1000.00', '2023-09-09 15:35:02', '2023-09-09 15:35:02'),
(4, 8, NULL, 3, '1000.00', '2023-09-09 15:35:02', '2023-09-09 15:35:02'),
(4, 8, NULL, 2, '1000.00', '2023-09-09 15:35:02', '2023-09-09 15:35:02'),
(4, 7, NULL, 1, '1000.00', '2023-09-09 15:35:13', '2023-09-09 15:35:13'),
(4, 8, NULL, 2, '1000.00', '2023-09-09 15:35:13', '2023-09-09 15:35:13'),
(4, 6, NULL, 2, '1000.00', '2023-09-09 15:35:13', '2023-09-09 15:35:13'),
(5, 7, NULL, 1, '1000.00', '2023-09-09 15:35:40', '2023-09-09 15:35:40'),
(5, 4, NULL, 2, '1000.00', '2023-09-09 15:35:40', '2023-09-09 15:35:40'),
(5, 5, NULL, 2, '1000.00', '2023-09-09 15:35:40', '2023-09-09 15:35:40'),
(6, 4, NULL, 4, '1000.00', '2023-09-09 15:35:57', '2023-09-09 15:35:57'),
(6, 3, NULL, 6, '1000.00', '2023-09-09 15:35:57', '2023-09-09 15:35:57'),
(6, 2, NULL, 2, '1000.00', '2023-09-09 15:35:57', '2023-09-09 15:35:57'),
(7, 4, NULL, 1, '1000.00', '2023-09-09 15:36:12', '2023-09-09 15:36:12'),
(7, 3, NULL, 1, '1000.00', '2023-09-09 15:36:12', '2023-09-09 15:36:12'),
(7, 2, NULL, 1, '1000.00', '2023-09-09 15:36:12', '2023-09-09 15:36:12'),
(8, 3, NULL, 2, '1000.00', '2023-09-09 15:36:26', '2023-09-09 15:36:26'),
(8, 2, NULL, 1, '1000.00', '2023-09-09 15:36:26', '2023-09-09 15:36:26'),
(8, 1, NULL, 3, '1000.00', '2023-09-09 15:36:26', '2023-09-09 15:36:26'),
(9, 3, NULL, 2, '1000.00', '2023-09-09 15:36:37', '2023-09-09 15:36:37'),
(9, 2, NULL, 1, '1000.00', '2023-09-09 15:36:37', '2023-09-09 15:36:37'),
(9, 2, NULL, 1, '1000.00', '2023-09-09 15:36:37', '2023-09-09 15:36:37'),
(10, 3, NULL, 2, '1000.00', '2023-09-09 15:36:55', '2023-09-09 15:36:55'),
(10, 1, NULL, 3, '1000.00', '2023-09-09 15:36:55', '2023-09-09 15:54:08'),
(10, 3, NULL, 2, '1000.00', '2023-09-09 15:36:55', '2023-09-09 15:36:55'),
(11, 5, NULL, 2, '1000.00', '2023-09-09 15:37:09', '2023-09-09 15:37:09'),
(11, 3, NULL, 1, '1000.00', '2023-09-09 15:37:09', '2023-09-09 15:37:09'),
(11, 1, NULL, 3, '1000.00', '2023-09-09 15:37:09', '2023-09-09 15:54:08'),
(12, 3, NULL, 5, '1000.00', '2023-09-09 15:37:27', '2023-09-09 15:37:27'),
(12, 3, NULL, 1, '1000.00', '2023-09-09 15:37:27', '2023-09-09 15:37:27'),
(12, 3, NULL, 2, '1000.00', '2023-09-09 15:37:27', '2023-09-09 15:37:27'),
(13, 1, NULL, 3, '1000.00', '2023-09-09 15:37:40', '2023-09-09 15:54:08'),
(13, 3, NULL, 1, '1000.00', '2023-09-09 15:37:40', '2023-09-09 15:37:40'),
(13, 5, NULL, 3, '1000.00', '2023-09-09 15:37:40', '2023-09-09 15:37:40'),
(14, 4, NULL, 5, '1000.00', '2023-09-09 15:37:52', '2023-09-09 15:37:52'),
(14, 3, NULL, 1, '1000.00', '2023-09-09 15:37:52', '2023-09-09 15:37:52'),
(14, 5, NULL, 3, '1000.00', '2023-09-09 15:37:52', '2023-09-09 15:37:52'),
(15, 4, NULL, 5, '1000.00', '2023-09-09 15:38:09', '2023-09-09 15:38:09'),
(15, 4, NULL, 4, '1000.00', '2023-09-09 15:38:09', '2023-09-09 15:38:09'),
(15, 4, NULL, 3, '1000.00', '2023-09-09 15:38:09', '2023-09-09 15:38:09');

--
-- Triggers `transaction_details`
--
DELIMITER $$
CREATE TRIGGER `after_transaction_details_delete` AFTER DELETE ON `transaction_details` FOR EACH ROW BEGIN
    DECLARE transaction_id_to_update INT;
    DECLARE total_quantity INT;

    SET transaction_id_to_update = OLD.transaction_id;

    SELECT SUM(qty) INTO total_quantity
    FROM transaction_details
    WHERE transaction_id = transaction_id_to_update;

    UPDATE transaction
    SET total_qty = total_quantity
    WHERE id = transaction_id_to_update;
END
$$
DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `status` smallint(6) DEFAULT NULL,
  `dob` date DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `status`, `dob`, `gender`, `created_at`, `updated_at`) VALUES
(1, 'nasywa', NULL, '2004-04-05', 'P', '2023-09-09 14:49:35', '2023-09-09 14:49:35'),
(2, 'dinda', NULL, '2003-02-03', 'P', '2023-09-09 14:49:35', '2023-09-09 14:49:35'),
(3, 'nafa', NULL, '2003-01-23', 'P', '2023-09-09 14:49:35', '2023-09-09 14:49:35'),
(4, 'khansa', NULL, '2002-09-13', 'L', '2023-09-09 14:49:35', '2023-09-09 14:49:35'),
(5, 'luthfan', NULL, '2003-01-09', 'L', '2023-09-09 14:49:35', '2023-09-09 14:49:35');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `operators`
--
ALTER TABLE `operators`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `operator_id` (`operator_id`),
  ADD KEY `product_type_id` (`product_type_id`);

--
-- Indexes for table `product_description`
--
ALTER TABLE `product_description`
  ADD PRIMARY KEY (`id`),
  ADD KEY `product_id` (`product_id`);

--
-- Indexes for table `product_types`
--
ALTER TABLE `product_types`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `fk_payment` (`payment_method_id`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD KEY `transaction_id` (`transaction_id`),
  ADD KEY `fk_productId` (`product_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`operator_id`) REFERENCES `operators` (`id`),
  ADD CONSTRAINT `products_ibfk_3` FOREIGN KEY (`product_type_id`) REFERENCES `product_types` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `fk_payment` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_methods` (`id`),
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`transaction_id`) REFERENCES `transactions` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
