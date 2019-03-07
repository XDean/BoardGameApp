package xdean.mini.boardgame.server.security.handler;

import org.springframework.data.jpa.repository.JpaRepository;

import xdean.mini.boardgame.server.security.model.AccessTokenEntity;

public interface AccessTokenRepo extends JpaRepository<AccessTokenEntity, String> {

}
