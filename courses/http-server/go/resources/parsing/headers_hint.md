# Hint: Implementar parseHeaders

## Paso 1: Crear el mapa e iniciar el loop

```go
headers := make(map[string]string)
for {
    line, err := reader.ReadString('\n')
    if err != nil {
        return nil, fmt.Errorf("error leyendo headers: %w", err)
    }
    line = strings.TrimRight(line, "\r\n")
```

## Paso 2: Detectar fin de headers

```go
    if line == "" {
        break  // Línea vacía = fin de headers (\r\n\r\n)
    }
```

## Paso 3: Parsear cada header

```go
    parts := strings.SplitN(line, ":", 2)
    if len(parts) != 2 {
        continue  // Header malformado, ignorar
    }
    key := strings.TrimSpace(parts[0])
    val := strings.TrimSpace(parts[1])
    headers[key] = val
}
return headers, nil
```

Usamos `SplitN` con límite 2 para manejar valores que contienen `:` (como URLs).
