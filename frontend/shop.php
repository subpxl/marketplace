<!-- product.php -->
<?php
include_once 'header.php';  // Include the header

?>

<section>
    <!-- Banner -->

    <section>
<!-- Banner -->
<div class="container mx-auto px-4 py-8">
    <div class="carousel w-full max-w-4xl mx-auto rounded-box shadow-lg" id="banner-carousel">
        <div id="slide1" class="carousel-item relative w-full">
            <img src="/static/1.jpg" class="w-full object-cover h-[300px] md:h-[400px]" alt="Banner 1" />
            <div class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
                <button onclick="moveSlide(-1)" class="btn btn-circle">❮</button> 
                <button onclick="moveSlide(1)" class="btn btn-circle">❯</button>
            </div>
        </div> 
        <div id="slide2" class="carousel-item relative w-full">
            <img src="/static/2.jpg" class="w-full object-cover h-[300px] md:h-[400px]" alt="Banner 2" />
            <div class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
                <button onclick="moveSlide(-1)" class="btn btn-circle">❮</button> 
                <button onclick="moveSlide(1)" class="btn btn-circle">❯</button>
            </div>
        </div> 
        <div id="slide3" class="carousel-item relative w-full">
            <img src="/static/3.jpg" class="w-full object-cover h-[300px] md:h-[400px]" alt="Banner 3" />
            <div class="absolute flex justify-between transform -translate-y-1/2 left-5 right-5 top-1/2">
                <button onclick="moveSlide(-1)" class="btn btn-circle">❮</button> 
                <button onclick="moveSlide(1)" class="btn btn-circle">❯</button>
            </div>
        </div>
    </div>
</div>

<script>
    let currentSlide = 1;
    const totalSlides = 3;
    
    function moveSlide(direction) {
        currentSlide += direction;
        if (currentSlide > totalSlides) {
            currentSlide = 1;
        } else if (currentSlide < 1) {
            currentSlide = totalSlides;
        }
        document.querySelector(`#slide${currentSlide}`).scrollIntoView({
            behavior: 'smooth',
            block: 'nearest',
            inline: 'start'
        });
    }
    
    // Prevent default anchor behavior
    document.querySelectorAll('#banner-carousel .btn-circle').forEach(btn => {
        btn.addEventListener('click', (e) => {
            e.preventDefault();
        });
    });
    </script>

</section>
    <!-- Sale Text -->
    <div class="bg-accent text-accent-content text-center py-2">
        <p class="text-lg font-bold">Sale! Up to 50% off on selected items!</p>
    </div>

    <!-- Product Grid -->
     
    <section class="container mx-auto ">
        <h2 class="text-2xl font-bold mb-4 ">Our Products</h2>
        <?php include 'product_grid.php'; ?>
    </section>


<!-- Featured Products -->
<section class="container mx-auto ">
    <h2 class="text-2xl font-bold mb-4">Featured Products</h2>
    <?php include 'product_grid.php'; ?>
        
</section>

    <!-- About Us -->
    <section class="container mx-auto px-4 py-8">
        <h2 class="text-2xl font-bold mb-4">About Us</h2>
        <p class="mb-4">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam in dui mauris. Vivamus hendrerit arcu sed erat molestie vehicula. Sed auctor neque eu tellus rhoncus ut eleifend nibh porttitor. Ut in nulla enim. Phasellus molestie magna non est bibendum non venenatis nisl tempor. Suspendisse dictum feugiat nisl ut dapibus. Mauris iaculis porttitor posuere. Praesent id metus massa, ut blandit odio. Proin quis tortor orci. Etiam at risus et justo dignissim congue.</p>
    </section>

    <!-- Google Map -->
    <section class="container mx-auto px-4 py-8">
        <h2 class="text-2xl font-bold mb-4">Our Location</h2>
        <div class="w-full h-64">
            <!-- Replace with actual Google Maps embed code -->
            <iframe src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d3022.1422937950147!2d-73.98731968482413!3d40.75889497932681!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x89c25855c6480299%3A0x55194ec5a1ae072e!2sTimes+Square!5e0!3m2!1sen!2sus!4v1560412335569!5m2!1sen!2sus" width="100%" height="100%" frameborder="0" style="border:0" allowfullscreen></iframe>
        </div>
    </section>


    <!-- Contact Details and Form -->
<section class="container mx-auto px-4 py-8 bg-white">
    <h2 class="text-2xl font-bold mb-4">Contact Us</h2>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <div>
            <p>Email: info@example.com</p>
            <p>Phone: +1 (123) 456-7890</p>
            <p>Address: 123 Main St, City, Country</p>
        </div>
        <form id="contactForm" onsubmit="sendToWhatsApp(event)">
            <input type="text" id="name" placeholder="Name" class="input input-bordered w-full mb-2" required>
            <input type="email" id="email" placeholder="Email (optional)" class="input input-bordered w-full mb-2">
            <textarea id="message" class="textarea textarea-bordered w-full mb-2" placeholder="Message" required></textarea>
            <button type="submit" class="btn btn-primary">Send Message</button>
        </form>
    </div>
</section>
</section>


<?php
include_once 'footer.php';  // Include the footer
