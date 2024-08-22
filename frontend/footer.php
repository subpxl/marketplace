<!-- footer.php -->
<script>
function sendToWhatsApp(event) {
    event.preventDefault();
    
    var name = document.getElementById('name').value;
    var email = document.getElementById('email').value;
    var message = document.getElementById('message').value;
    
    var whatsappNumber = "1234567890"; // Replace with the seller's WhatsApp number
    var whatsappMessage = "Name: " + name + "%0A";
    
    if (email) {
        whatsappMessage += "Email: " + email + "%0A";
    }
    
    whatsappMessage += "Message: " + message;
    
    var whatsappURL = "https://wa.me/" + whatsappNumber + "?text=" + encodeURIComponent(whatsappMessage);
    
    window.open(whatsappURL, '_blank');
}
</script>
<!-- Footer -->
<footer class="bg-neutral text-neutral-content">
    <div class="container mx-auto px-4 py-8">
        <div class="flex flex-col items-center">
            <!-- Column 1: Logo -->
            <div class="w-full text-center mb-6">
                <h3 class="text-lg font-bold mb-2">Shop Name</h3>
                <!-- <p>Providing reliable products since 2023</p> -->
            </div>

            <!-- Columns for Links and More Links -->
            <div class="w-full flex justify-center mb-6">
                <!-- Column 2: Links Column 1 -->
                <div class="w-1/2 px-2">
                    <h3 class="text-lg font-bold mb-2 text-center">Links</h3>
                    <ul class="text-center">
                        <li><a href="#" class="link link-hover">Terms of Service</a></li>
                        <li><a href="#" class="link link-hover">Privacy Policy</a></li>
                    </ul>
                </div>
                
                <!-- Column 3: Links Column 2 -->
                <div class="w-1/2 px-2">
                    <h3 class="text-lg font-bold mb-2 text-center">More Links</h3>
                    <ul class="text-center">
                        <li><a href="#" class="link link-hover">About</a></li>
                        <li><a href="#" class="link link-hover">Contact</a></li>
                    </ul>
                </div>
            </div>

            <!-- Column 4: Social Media Links -->
            <div class="w-full flex justify-center gap-4 mb-6">
                <a href="#" class="btn btn-ghost btn-square">
                    <i class="fab fa-facebook-f text-2xl"></i>
                </a>
                <a href="#" class="btn btn-ghost btn-square">
                    <i class="fab fa-youtube text-2xl"></i>
                </a>
                <a href="#" class="btn btn-ghost btn-square">
                    <i class="fab fa-twitter text-2xl"></i>
                </a>
                <a href="#" class="btn btn-ghost btn-square">
                    <i class="fab fa-instagram text-2xl"></i>
                </a>
                <a href="#" class="btn btn-ghost btn-square">
                    <i class="fab fa-linkedin-in text-2xl"></i>
                </a>
            </div>
        </div>
        
        <!-- Bottom Tag -->
        <div class="text-center mt-4">
            <p class="text-sm">
                Made with love by <a href="https://digitalbazaar.io" class="link link-hover">digitalbazaar.io</a>
            </p>
        </div>
    </div>
</footer>
    <!-- Mobile Bottom Navigation -->
    <div class="btm-nav mobile-view bg-primary ">
        <button class="text-primary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6" /></svg>
            <span class="btm-nav-label text-white">Home</span>
        </button>
        <button class="text-primary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5  text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
            <span class="btm-nav-label  text-white">catalog</span>
        </button>
        <button class="text-primary">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5  text-white" fill="none" viewBox="0 0 24 24" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z" /></svg>
            <span class="btm-nav-label  text-white">Cart</span>
        </button>

    </div>
        <!-- Floating WhatsApp Button -->
        <a href="https://wa.me/919109517579?text=I'm%20interested%20in%20your%20car%20for%20sale" target="_blank" class="fixed bottom-20 right-4 bg-green-500 text-white rounded-full p-3 shadow-lg md:bottom-4">
            <i class="fab fa-whatsapp fa-2x"></i>
        </a>
    


        <script>
            function shareStore() {
    if (navigator.share) {
        navigator.share({
            title: 'Check out this awesome store!',
            text: 'I found this great online store. You should check it out!',
            url: window.location.href
        }).then(() => {
            console.log('Thanks for sharing!');
        })
        .catch(console.error);
    } else {
        // Fallback for browsers that don't support the Web Share API
        alert('Share this page: ' + window.location.href);
    }
}
        </script>


</body>
</html>