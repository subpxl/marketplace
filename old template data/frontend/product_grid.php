<div class="container mx-auto p-4">
            <div class="grid grid-cols-2 gap-4">
            <?php
    for ($i = 0; $i < 4; $i++) {
        echo '<div class="card bg-base-100 shadow-lg hover:shadow-xl transition-shadow duration-300 overflow-hidden">
                    <figure class="px-4 pt-4">
                        <img src="/static/'. ($i + 1) .'.jpg" alt="Product" class="rounded-xl h-48 w-full object-cover transform hover:scale-105 transition-transform duration-300" />
                    </figure>
                    <div class="card-body p-4">
                        <h2 class="card-title text-lg font-bold">Elegant Tote</h2>
                        <p class="text-sm text-gray-600 mb-2">Luxurious leather bag</p>
                        <p class="text-primary font-semibold text-lg mb-2">$199.99</p>
                        <div class="card-actions">
                            <button class="btn btn-primary btn-sm w-full text-white">Buy Now</button>
                        </div>
                    </div>
                </div>'; 
    }
    ?>         
            </div>
        </div>