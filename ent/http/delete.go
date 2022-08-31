// Code generated by entc, DO NOT EDIT.

package http

import (
	"net/http"
	"strconv"
	"t/ent"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// Delete removes a ent.Compartment from the database.
func (h CompartmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Delete"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	err = h.client.Compartment.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		default:
			l.Error("could-not-delete-compartment", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("compartment deleted", zap.Int("id", id))
	w.WriteHeader(http.StatusNoContent)
}

// Delete removes a ent.Entry from the database.
func (h EntryHandler) Delete(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Delete"))
	// ID is URL parameter.
	var err error
	id := chi.URLParam(r, "id")
	err = h.client.Entry.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.String("id", id))
			NotFound(w, msg)
		default:
			l.Error("could-not-delete-entry", zap.Error(err), zap.String("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("entry deleted", zap.String("id", id))
	w.WriteHeader(http.StatusNoContent)
}

// Delete removes a ent.Fridge from the database.
func (h FridgeHandler) Delete(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Delete"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	err = h.client.Fridge.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		default:
			l.Error("could-not-delete-fridge", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("fridge deleted", zap.Int("id", id))
	w.WriteHeader(http.StatusNoContent)
}

// Delete removes a ent.Item from the database.
func (h ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {
	l := h.log.With(zap.String("method", "Delete"))
	// ID is URL parameter.
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		l.Error("error getting id from url parameter", zap.String("id", chi.URLParam(r, "id")), zap.Error(err))
		BadRequest(w, "id must be an integer")
		return
	}
	err = h.client.Item.DeleteOneID(id).Exec(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			msg := stripEntError(err)
			l.Info(msg, zap.Error(err), zap.Int("id", id))
			NotFound(w, msg)
		default:
			l.Error("could-not-delete-item", zap.Error(err), zap.Int("id", id))
			InternalServerError(w, nil)
		}
		return
	}
	l.Info("item deleted", zap.Int("id", id))
	w.WriteHeader(http.StatusNoContent)
}