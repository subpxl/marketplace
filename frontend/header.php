<!-- header.php -->


<!DOCTYPE html>
<html lang="en" data-theme="light">
<head>

    <meta charset="UTF-8">
        <!-- <meta name="viewport" content="width=device-width, initial-scale=0.5"> -->
        <meta name="viewport"  content="width=375, initial-scale=1.0, user-scalable=no">

    <title>Ecommerce Shop</title>
    <link href="https://cdn.jsdelivr.net/npm/daisyui@3.1.0/dist/full.css" rel="stylesheet" type="text/css" />
    <script src="https://cdn.tailwindcss.com"></script>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css">
</head>
<style>
 .preloader {
            position: fixed;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background-color: #fff;
            display: flex;
            justify-content: center;
            align-items: center;
            z-index: 9999;
        }

        .mobile-view {
        max-width: 375px;
        margin: 0 auto;
        border: 1px solid #ddd;
    }

    .btn-primary {
            background-color:  #00234E;
            border: none;
        }

        .bg-primary {
            background-color:  #00234E;
            border: none;
        }
        

</style>
<body class="mobile-view">

        <!-- Preloader -->

        <!-- Black Banner Ribbon -->
        <div class="bg-primary text-white text-center py-2">
            <p>Latest store sale: 50% off on all summer collections!</p>
        </div>
    

    <!-- Sticky Navbar -->
    <header class="bg-white shadow-lg sticky top-0 z-50">
        <div class="container mx-auto p-4 flex justify-between items-center">
            <h1 class="text-2xl font-bold text-black">Shop Name</h1>
            <!-- Share icon for mobile -->
            <div class="md">
                <button onclick="shareStore()" class="btn btn-square btn-ghost">
                    <i class="fas fa-share-alt"></i>
                </button>
            </div>
            
        </div>
    </header>

