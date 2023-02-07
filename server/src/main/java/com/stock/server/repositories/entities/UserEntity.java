package com.stock.server.repositories.entities;

import java.util.UUID;

public record UserEntity(UUID userUuid, String email, String username, String password) {
}
