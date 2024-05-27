-- CreateTable
CREATE TABLE
    `Account` (
        `id` INTEGER NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(191) NOT NULL,
        `password` VARCHAR(191) NOT NULL,
        `phone` VARCHAR(191) NULL,
        `location` VARCHAR(191) NULL,
        `email` VARCHAR(191) NOT NULL,
        `role` ENUM ('ADMIN', 'STORE', 'CUSTOMER') NOT NULL DEFAULT 'CUSTOMER',
        `status` BOOLEAN NOT NULL DEFAULT true,
        `imageAvatar` VARCHAR(191) NOT NULL DEFAULT 'https://www.nasco.co.th/wp-content/uploads/2022/06/placeholder.png',
        `storeName` VARCHAR(191) NULL,
        `storeLocation` VARCHAR(191) NULL,
        `permissionId` INTEGER NULL,
        `createdAt` DATETIME (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
        `updatedAt` DATETIME (3) NOT NULL,
        UNIQUE INDEX `Account_email_key` (`email`),
        UNIQUE INDEX `Account_permissionId_key` (`permissionId`),
        INDEX `Account_email_idx` (`email`),
        PRIMARY KEY (`id`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE
    `Product` (
        `id` INTEGER NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(191) NOT NULL,
        `price` DECIMAL(65, 30) NOT NULL,
        `stock` INTEGER NOT NULL,
        `detail` VARCHAR(191) NULL,
        `status` BOOLEAN NOT NULL,
        `categoryId` INTEGER NOT NULL,
        `storeId` INTEGER NOT NULL,
        `productAvatar` VARCHAR(191) NOT NULL DEFAULT 'https://www.nasco.co.th/wp-content/uploads/2022/06/placeholder.png',
        `createdAt` DATETIME (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
        `updatedAt` DATETIME (3) NOT NULL,
        PRIMARY KEY (`id`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE
    `ProductCategory` (
        `id` INTEGER NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(191) NOT NULL,
        `status` BOOLEAN NOT NULL,
        `code` VARCHAR(191) NULL,
        `detail` VARCHAR(191) NULL,
        `createdAt` DATETIME (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
        `updatedAt` DATETIME (3) NOT NULL,
        PRIMARY KEY (`id`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE
    `Order` (
        `id` INTEGER NOT NULL AUTO_INCREMENT,
        `orderId` VARCHAR(191) NOT NULL,
        `totalAmount` DECIMAL(65, 30) NOT NULL,
        `topic` VARCHAR(191) NULL,
        `sumPrice` DOUBLE NOT NULL,
        `state` ENUM (
            'PENDING',
            'PREPARED',
            'SEND',
            'SUCCEED',
            'REJECTED'
        ) NOT NULL DEFAULT 'PENDING',
        `deliveryType` VARCHAR(191) NULL,
        `parcelNumber` VARCHAR(191) NULL,
        `sentDate` DATETIME (3) NULL,
        `customerId` INTEGER NOT NULL,
        `storeId` INTEGER NOT NULL,
        `bankId` INTEGER NOT NULL,
        `createdAt` DATETIME (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
        `updatedAt` DATETIME (3) NOT NULL,
        UNIQUE INDEX `Order_orderId_key` (`orderId`),
        PRIMARY KEY (`id`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE
    `OrderProduct` (
        `orderId` INTEGER NOT NULL,
        `productId` INTEGER NOT NULL,
        `quantity` INTEGER NOT NULL,
        `createdAt` DATETIME (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
        `updatedAt` DATETIME (3) NOT NULL,
        PRIMARY KEY (`orderId`, `productId`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE
    `Bank` (
        `id` INTEGER NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(191) NOT NULL,
        `accNumber` VARCHAR(191) NOT NULL,
        `accName` VARCHAR(191) NOT NULL,
        `status` VARCHAR(191) NOT NULL,
        `avatarUrl` VARCHAR(191) NOT NULL DEFAULT 'https://www.nasco.co.th/wp-content/uploads/2022/06/placeholder.png',
        `storeId` INTEGER NOT NULL,
        `createdAt` DATETIME (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
        `updatedAt` DATETIME (3) NOT NULL,
        PRIMARY KEY (`id`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE
    `AdminPermission` (
        `id` INTEGER NOT NULL AUTO_INCREMENT,
        `menuPermission` JSON NOT NULL,
        PRIMARY KEY (`id`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE
    `File` (
        `id` INTEGER NOT NULL AUTO_INCREMENT,
        `name` VARCHAR(191) NOT NULL,
        `pathUrl` VARCHAR(191) NOT NULL,
        `type` VARCHAR(191) NOT NULL,
        `entityType` ENUM ('ACCOUNT', 'ORDER', 'PRODUCT', 'BANK') NOT NULL,
        `entityId` INTEGER NOT NULL,
        `accountId` INTEGER NULL,
        `orderId` INTEGER NULL,
        `createdAt` DATETIME (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
        `updatedAt` DATETIME (3) NOT NULL,
        PRIMARY KEY (`id`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE
    `Log` (
        `id` INTEGER NOT NULL AUTO_INCREMENT,
        `fullName` VARCHAR(191) NOT NULL,
        `menuRequest` VARCHAR(191) NOT NULL,
        `actionRequest` VARCHAR(191) NOT NULL,
        `createdAt` DATETIME (3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3),
        `updatedAt` DATETIME (3) NOT NULL,
        `accountId` INTEGER NOT NULL,
        PRIMARY KEY (`id`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- CreateTable
CREATE TABLE
    `_StoreProductCategory` (
        `A` INTEGER NOT NULL,
        `B` INTEGER NOT NULL,
        UNIQUE INDEX `_StoreProductCategory_AB_unique` (`A`, `B`),
        INDEX `_StoreProductCategory_B_index` (`B`)
    ) DEFAULT CHARACTER
SET
    utf8mb4 COLLATE utf8mb4_unicode_ci;

-- AddForeignKey
ALTER TABLE `Account` ADD CONSTRAINT `Account_permissionId_fkey` FOREIGN KEY (`permissionId`) REFERENCES `AdminPermission` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `Product` ADD CONSTRAINT `Product_categoryId_fkey` FOREIGN KEY (`categoryId`) REFERENCES `ProductCategory` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `Product` ADD CONSTRAINT `Product_storeId_fkey` FOREIGN KEY (`storeId`) REFERENCES `Account` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `Order` ADD CONSTRAINT `Order_customerId_fkey` FOREIGN KEY (`customerId`) REFERENCES `Account` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `Order` ADD CONSTRAINT `Order_storeId_fkey` FOREIGN KEY (`storeId`) REFERENCES `Account` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `Order` ADD CONSTRAINT `Order_bankId_fkey` FOREIGN KEY (`bankId`) REFERENCES `Bank` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `OrderProduct` ADD CONSTRAINT `OrderProduct_orderId_fkey` FOREIGN KEY (`orderId`) REFERENCES `Order` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `OrderProduct` ADD CONSTRAINT `OrderProduct_productId_fkey` FOREIGN KEY (`productId`) REFERENCES `Product` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `Bank` ADD CONSTRAINT `Bank_storeId_fkey` FOREIGN KEY (`storeId`) REFERENCES `Account` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `File` ADD CONSTRAINT `AccountFileRelation` FOREIGN KEY (`entityId`) REFERENCES `Account` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `File` ADD CONSTRAINT `OrderFileRelation` FOREIGN KEY (`entityId`) REFERENCES `Order` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `File` ADD CONSTRAINT `ProductFileRelation` FOREIGN KEY (`entityId`) REFERENCES `Product` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `File` ADD CONSTRAINT `BankFileRelation` FOREIGN KEY (`entityId`) REFERENCES `Bank` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `File` ADD CONSTRAINT `File_orderId_fkey` FOREIGN KEY (`orderId`) REFERENCES `Order` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `Log` ADD CONSTRAINT `Log_accountId_fkey` FOREIGN KEY (`accountId`) REFERENCES `Account` (`id`) ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `_StoreProductCategory` ADD CONSTRAINT `_StoreProductCategory_A_fkey` FOREIGN KEY (`A`) REFERENCES `Account` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE `_StoreProductCategory` ADD CONSTRAINT `_StoreProductCategory_B_fkey` FOREIGN KEY (`B`) REFERENCES `ProductCategory` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;