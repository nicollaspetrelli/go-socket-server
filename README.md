# GoLang WebSocket Server
Realtime Locations example with GoLang and Socket.IO

## Annotations

No Socket.IO, uma sala é simplesmente um grupo de sockets (conexões) que você pode transmitir mensagens. Quando um socket se conecta, você pode adicioná-lo a uma sala específica.

Para este caso de uso onde enviamos as localizações de dispositivos IOTs de várias empresas, usar salas (rooms) seria mais adequado. Aqui estão algumas razões:

1. **Isolamento**: As salas permitem isolar as mensagens de diferentes empresas. Cada empresa teria sua própria sala, e as mensagens enviadas a uma sala só seriam recebidas pelos sockets nessa sala.

2. **Flexibilidade**: Usar salas permite que você envie diferentes tipos de eventos para a mesma sala. Por exemplo, você poderia ter eventos `location`, `status`, `update`, etc., todos enviados para a mesma sala.

3. **Escalabilidade**: À medida que o número de empresas aumenta, o uso de salas torna-se mais escalável. Você pode facilmente adicionar novas salas para novas empresas sem ter que mudar a lógica do seu código.

Portanto, para este caso de uso, usar salas seria a abordagem mais adequada.