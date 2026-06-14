Project Overview: PokeTestAPI CLI & Cache Layer

This Go-based application serves as a high-performance, command-line interface (CLI) and caching server designed to interface with an external Pokémon API. Acting as an intermediary middleware layer, the application performs CRUD-like operations to fetch, manage, and optimize Pokémon data retrieval before securely passing the processed information to downstream applications.
Key Technical Features

    Custom In-Memory Caching System: Integrates a thread-safe caching module (pokeCache) with an automated eviction policy (configured with a 7-second time-to-live base) to minimize external API latency, eliminate redundant network requests, and prevent API rate-limiting issues.

    Dynamic Command-Line Interface (CLI): Implements an interactive, loop-driven REPL (Read-Eval-Print Loop) environment using Go's native bufio.NewScanner to parse user inputs, sanitize text data, and map instructions efficiently.

    Extensible Command Architecture: Utilizes a decoupled callback configuration map (climap) to dynamically execute instructions and pass arguments, separating CLI input handling from core business and API logic.

    Downstream Pipeline Integration: Built to scale seamlessly as a lightweight data proxy, ensuring clean schema delivery of Pokémon metadata to interconnected microservices or frontend clients.
