<script>
    const handleLogin = async () => {
        const url = 'http://localhost:8080/auth/google';

        try {
            const response = await fetch(url, {
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            }); 

            if (!response.ok) {
                throw new Error('Failed to login! status: ' + response.status);
            }

            const data = await response.json();
            console.log(data);

            } catch (error) {
            console.error('Error during fetch:', error);
        }
    };
  </script>
  
  <div class="flex items-center justify-center h-screen bg-gray-100">
    <div class="bg-white p-8 rounded shadow-md w-96">
      <h2 class="text-2xl font-bold mb-6 text-center">Login</h2>
      <button 
        on:click={handleLogin} 
        class="w-full bg-blue-600 text-white py-2 rounded hover:bg-blue-700 transition duration-200">
        Login with Google
      </button>
    </div>
  </div>